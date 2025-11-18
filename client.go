package devin

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin/binding"

	"github.com/gassara-kys/go-devin/internal/httpclient"
	"github.com/gassara-kys/go-devin/pkg/knowledge"
	"github.com/gassara-kys/go-devin/pkg/playbooks"
	"github.com/gassara-kys/go-devin/pkg/secrets"
	"github.com/gassara-kys/go-devin/pkg/sessions"
)

const (
	defaultBaseURL       = "https://api.devin.ai/v1"
	defaultUserAgent     = "go-devin/0.1.0"
	defaultTimeout       = 30 * time.Second
	defaultRetryAttempts = 3
	defaultRetryInitial  = 500 * time.Millisecond
	defaultRetryMaxWait  = 2 * time.Second
)

type HTTPDoer = httpclient.Doer
type RetryConfig = httpclient.RetryConfig

type Client struct {
	cfg       httpclient.Config
	transport httpclient.Transport
	validator binding.StructValidator

	logger *slog.Logger

	Sessions  *sessions.Service
	Secrets   *secrets.Service
	Knowledge *knowledge.Service
	Playbooks *playbooks.Service
}

type Option func(*Client)

func NewClient(apiKey string, opts ...Option) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("api key is required")
	}

	logger := slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{}))
	cfg := httpclient.Config{
		BaseURL:   defaultBaseURL,
		APIKey:    apiKey,
		UserAgent: defaultUserAgent,
		HTTPClient: &http.Client{
			Timeout: defaultTimeout,
		},
		Logger: logger,
		Retry: RetryConfig{
			MaxAttempts:    defaultRetryAttempts,
			InitialBackoff: defaultRetryInitial,
			MaxBackoff:     defaultRetryMaxWait,
		},
	}

	c := &Client{
		cfg:       cfg,
		logger:    logger,
		validator: binding.Validator,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.cfg.HTTPClient == nil {
		c.cfg.HTTPClient = &http.Client{Timeout: defaultTimeout}
	}
	c.cfg.Logger = c.logger

	exec := httpclient.NewExecutor(&c.cfg)
	c.transport = exec

	validate := func(v any) error {
		if c.validator == nil || v == nil {
			return nil
		}
		return c.validator.ValidateStruct(v)
	}

	c.Sessions = sessions.NewService(exec, validate)
	c.Secrets = secrets.NewService(exec, validate)
	c.Knowledge = knowledge.NewService(exec, validate)
	c.Playbooks = playbooks.NewService(exec, validate)

	return c, nil
}

func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		if baseURL != "" {
			c.cfg.BaseURL = baseURL
		}
	}
}

func WithHTTPClient(h HTTPDoer) Option {
	return func(c *Client) {
		if h != nil {
			c.cfg.HTTPClient = h
		}
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		if httpClient, ok := c.cfg.HTTPClient.(*http.Client); ok && timeout > 0 {
			httpClient.Timeout = timeout
		}
	}
}

func WithRetry(cfg RetryConfig) Option {
	return func(c *Client) {
		if cfg.MaxAttempts > 0 {
			c.cfg.Retry.MaxAttempts = cfg.MaxAttempts
		}
		if cfg.InitialBackoff > 0 {
			c.cfg.Retry.InitialBackoff = cfg.InitialBackoff
		}
		if cfg.MaxBackoff > 0 {
			c.cfg.Retry.MaxBackoff = cfg.MaxBackoff
		}
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(c *Client) {
		if logger != nil {
			c.logger = logger
			c.cfg.Logger = logger
		}
	}
}

func WithUserAgent(ua string) Option {
	return func(c *Client) {
		if ua != "" {
			c.cfg.UserAgent = ua
		}
	}
}
