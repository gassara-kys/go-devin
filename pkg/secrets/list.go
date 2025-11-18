package secrets

import (
	"context"
	"net/http"

	"github.com/gassara-kys/go-devin/pkg/types"
)

type ListResponse struct {
	Secrets []types.Secret `json:"secrets"`
}

func (s *Service) List(ctx context.Context) (*ListResponse, error) {
	var resp ListResponse
	if err := s.doJSON(ctx, http.MethodGet, "/secrets", nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
