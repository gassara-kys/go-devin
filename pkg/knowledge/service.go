package knowledge

import (
	"context"
	"net/url"

	"github.com/gassara-kys/go-devin/internal/httpclient"
)

type Service struct {
	transport httpclient.Transport
	validator func(any) error
}

func NewService(t httpclient.Transport, validate func(any) error) *Service {
	if t == nil {
		panic("knowledge: transport is required")
	}
	if validate == nil {
		validate = func(any) error { return nil }
	}
	return &Service{transport: t, validator: validate}
}

func (s *Service) doJSON(ctx context.Context, method, path string, query url.Values, payload any, out any) error {
	return s.transport.DoJSON(ctx, method, path, query, payload, out)
}

func (s *Service) validateStruct(v any) error {
	if s.validator == nil || v == nil {
		return nil
	}
	return s.validator(v)
}
