package playbooks

import (
	"context"
	"net/http"

	"github.com/gassara-kys/go-devin/pkg/types"
)

// CreateRequest carries the payload for creating a playbook.
type CreateRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
	Macro string `json:"macro,omitempty"`
}

// CreateResponse contains the created playbook.
type CreateResponse struct {
	types.Playbook
}

// Create adds a new playbook.
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
