package playbooks

import (
	"context"
	"net/http"

	"github.com/gassara-kys/go-devin/pkg/types"
)

type CreateRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
	Macro string `json:"macro,omitempty"`
}

type CreateResponse struct {
	types.Playbook
}

func (s *Service) Create(ctx context.Context, req CreateRequest) (*CreateResponse, error) {
	if err := s.validateStruct(req); err != nil {
		return nil, err
	}
	var resp CreateResponse
	if err := s.doJSON(ctx, http.MethodPost, "/playbooks", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
