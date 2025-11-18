package testutil

import (
	"io"
	"log/slog"
	"net/http"
	"testing"
	"time"

	"github.com/gassara-kys/go-devin/internal/httpclient"
	"github.com/gassara-kys/go-devin/internal/testtransport"
)

// NewExecutor builds an httpclient.Transport backed by the provided RoundTripFunc.
func NewExecutor(t *testing.T, fn testtransport.RoundTripFunc) httpclient.Transport {
	t.Helper()
	cfg := httpclient.Config{
		BaseURL:   "https://api.test",
		APIKey:    "token",
		UserAgent: "test-agent",
		HTTPClient: &http.Client{
			Transport: fn,
			Timeout:   30 * time.Second,
		},
		Logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
		Retry: httpclient.RetryConfig{
			MaxAttempts:    1,
			InitialBackoff: 10 * time.Millisecond,
			MaxBackoff:     20 * time.Millisecond,
		},
	}
	return httpclient.NewExecutor(&cfg)
}
