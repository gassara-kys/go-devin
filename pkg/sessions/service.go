package sessions

import (
	"context"
	"net/url"

	"github.com/gassara-kys/go-devin/internal/httpclient"
)

// Service exposes the sessions-related API endpoints.
type Service struct {
	transport httpclient.Transport
	validate  func(any) error
}

// NewService builds a sessions Service instance.
func NewService(t httpclient.Transport, validate func(any) error) *Service {
	if t == nil {
		panic("sessions: transport is required")
	}
	if validate == nil {
		validate = func(any) error { return nil }
	}
	return &Service{transport: t, validate: validate}
}

func (s *Service) validateStruct(payload any) error {
	if s.validate == nil || payload == nil {
		return nil
	}
	return s.validate(payload)
}

func (s *Service) doJSON(ctx context.Context, method, path string, query url.Values, payload any, out any) error {
	return s.transport.DoJSON(ctx, method, path, query, payload, out)
}

func (s *Service) doBytes(ctx context.Context, method, path string, query url.Values, body []byte, contentType, accept string) ([]byte, error) {
	return s.transport.DoBytes(ctx, method, path, query, body, contentType, accept)
}
