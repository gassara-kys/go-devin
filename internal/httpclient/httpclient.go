package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Doer represents the minimal interface required from *http.Client.
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// RetryConfig tunes the exponential backoff retry behavior.
type RetryConfig struct {
	MaxAttempts    int
	InitialBackoff time.Duration
	MaxBackoff     time.Duration
}

// Config contains the information required to issue API calls.
type Config struct {
	BaseURL    string
	APIKey     string
	UserAgent  string
	HTTPClient Doer
	Logger     *slog.Logger
	Retry      RetryConfig
}

// Transport defines the public surface that higher level services consume.
type Transport interface {
	DoJSON(ctx context.Context, method, path string, query url.Values, payload any, out any) error
	DoBytes(ctx context.Context, method, path string, query url.Values, body []byte, contentType, accept string) ([]byte, error)
}

// Executor implements Transport on top of the configured HTTP client.
type Executor struct {
	cfg *Config
}

// NewExecutor builds an Executor using the provided Config.
func NewExecutor(cfg *Config) *Executor {
	return &Executor{cfg: cfg}
}

// DoJSON sends an HTTP request with optional JSON payload and decodes the JSON response.
func (e *Executor) DoJSON(ctx context.Context, method, p string, query url.Values, payload any, out any) error {
	var body []byte
	var err error
	if payload != nil {
		body, err = json.Marshal(payload)
		if err != nil {
			return err
		}
	}

	resp, err := e.doWithRetry(ctx, func() (*http.Request, error) {
		var reader io.Reader
		if body != nil {
			reader = bytes.NewReader(body)
		}
		req, err := e.newRequest(ctx, method, p, query, reader)
		if err != nil {
			return nil, err
		}
		if body != nil {
			req.Header.Set("Content-Type", "application/json")
		}
		req.Header.Set("Accept", "application/json")
		return req, nil
	})
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if out == nil || resp.StatusCode == http.StatusNoContent {
		_, err = io.Copy(io.Discard, resp.Body)
		return err
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(out); err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	return nil
}

// DoBytes sends an HTTP request with raw bytes and returns the raw response body.
func (e *Executor) DoBytes(ctx context.Context, method, p string, query url.Values, body []byte, contentType, accept string) ([]byte, error) {
	resp, err := e.doWithRetry(ctx, func() (*http.Request, error) {
		var reader io.Reader
		if body != nil {
			reader = bytes.NewReader(body)
		}
		req, err := e.newRequest(ctx, method, p, query, reader)
		if err != nil {
			return nil, err
		}
		if contentType != "" && body != nil {
			req.Header.Set("Content-Type", contentType)
		}
		if accept != "" {
			req.Header.Set("Accept", accept)
		}
		return req, nil
	})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	return io.ReadAll(resp.Body)
}

type requestBuilder func() (*http.Request, error)

func (e *Executor) newRequest(ctx context.Context, method, p string, query url.Values, body io.Reader) (*http.Request, error) {
	fullURL, err := e.resolveURL(p, query)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, fullURL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+e.cfg.APIKey)
	req.Header.Set("User-Agent", e.cfg.UserAgent)
	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/json")
	}
	return req, nil
}

func (e *Executor) resolveURL(p string, query url.Values) (string, error) {
	base, err := url.Parse(e.cfg.BaseURL)
	if err != nil {
		return "", fmt.Errorf("parse base url: %w", err)
	}

	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	rel := &url.URL{Path: p}
	if query != nil {
		rel.RawQuery = query.Encode()
	}
	return base.ResolveReference(rel).String(), nil
}

func (e *Executor) doWithRetry(ctx context.Context, builder requestBuilder) (*http.Response, error) {
	wait := e.cfg.Retry.InitialBackoff
	if wait <= 0 {
		wait = 200 * time.Millisecond
	}
	maxWait := e.cfg.Retry.MaxBackoff
	if maxWait <= 0 {
		maxWait = time.Second
	}
	if maxWait < wait {
		maxWait = wait
	}
	attempts := e.cfg.Retry.MaxAttempts
	if attempts < 1 {
		attempts = 1
	}

	var lastErr error
	for attempt := 1; attempt <= attempts; attempt++ {
		req, err := builder()
		if err != nil {
			return nil, err
		}

		if e.cfg.Logger != nil {
			e.cfg.Logger.DebugContext(ctx, "devin api request", "method", req.Method, "url", req.URL.String(), "attempt", attempt)
		}

		resp, err := e.cfg.HTTPClient.Do(req)
		if err != nil {
			lastErr = err
			if attempt == attempts || !shouldRetryError(err) {
				return nil, err
			}
			if err := sleep(ctx, wait); err != nil {
				return nil, err
			}
			wait = nextDuration(wait, maxWait)
			continue
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp, nil
		}

		body, readErr := io.ReadAll(resp.Body)
		closeErr := resp.Body.Close()
		if readErr != nil {
			return nil, readErr
		}
		if closeErr != nil {
			return nil, closeErr
		}

		apiErr := &APIError{
			StatusCode: resp.StatusCode,
			Body:       body,
			Detail:     parseAPIDetail(body),
		}
		lastErr = apiErr

		if attempt == attempts || !shouldRetryStatus(resp.StatusCode) {
			return nil, apiErr
		}

		if err := sleep(ctx, wait); err != nil {
			return nil, err
		}
		wait = nextDuration(wait, maxWait)
	}
	return nil, lastErr
}

func sleep(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

var retryableStatusCodes = map[int]struct{}{
	http.StatusTooManyRequests:     {},
	http.StatusRequestTimeout:      {},
	http.StatusTooEarly:            {},
	http.StatusInternalServerError: {},
	http.StatusBadGateway:          {},
	http.StatusServiceUnavailable:  {},
	http.StatusGatewayTimeout:      {},
}

func shouldRetryStatus(statusCode int) bool {
	_, ok := retryableStatusCodes[statusCode]
	return ok
}

func shouldRetryError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}
	var netErr net.Error
	return errors.As(err, &netErr)
}

func nextDuration(current, maxDuration time.Duration) time.Duration {
	next := current * 2
	if next > maxDuration {
		return maxDuration
	}
	return next
}

// APIError represents an error returned by the Devin API.
type APIError struct {
	StatusCode int
	Detail     string
	Body       []byte
}

func (e *APIError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Detail != "" {
		return fmt.Sprintf("devin api error: status=%d detail=%s", e.StatusCode, e.Detail)
	}
	return fmt.Sprintf("devin api error: status=%d", e.StatusCode)
}

func parseAPIDetail(body []byte) string {
	if len(body) == 0 {
		return ""
	}
	var payload map[string]any
	if err := json.Unmarshal(body, &payload); err == nil {
		for _, key := range []string{"detail", "message", "error"} {
			if v, ok := payload[key]; ok {
				if s, ok := v.(string); ok {
					return s
				}
			}
		}
	}
	return strings.TrimSpace(string(body))
}
