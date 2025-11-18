package secrets

import (
	"context"
	"net/http"
)

type CreateRequest struct {
	Type      string `json:"type" binding:"required"`
	Key       string `json:"key" binding:"required"`
	Value     string `json:"value" binding:"required"`
	Sensitive bool   `json:"sensitive"`
	Note      string `json:"note,omitempty"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

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
