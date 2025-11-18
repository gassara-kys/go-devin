package secrets

import (
	"context"
	"net/http"
)

// CreateRequest carries the payload for creating a secret.
type CreateRequest struct {
	Type      string `json:"type" binding:"required"`
	Key       string `json:"key" binding:"required"`
	Value     string `json:"value" binding:"required"`
	Sensitive bool   `json:"sensitive"`
	Note      string `json:"note,omitempty"`
}

// CreateResponse contains the identifier of the created secret.
type CreateResponse struct {
	ID string `json:"id"`
}

// Create stores a new secret.
func (s *Service) Create(ctx context.Context, req CreateRequest) (*CreateResponse, error) {
	if err := s.validateStruct(req); err != nil {
		return nil, err
	}
	var resp CreateResponse
	if err := s.doJSON(ctx, http.MethodPost, "/secrets", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
