package knowledge

import (
	"context"
	"net/http"

	"github.com/gassara-kys/go-devin/pkg/types"
)

type CreateRequest struct {
	Name               string `json:"name" binding:"required"`
	Body               string `json:"body" binding:"required"`
	TriggerDescription string `json:"trigger_description,omitempty"`
	ParentFolderID     string `json:"parent_folder_id,omitempty"`
	PinnedRepo         string `json:"pinned_repo,omitempty"`
}

type CreateResponse struct {
	types.Knowledge
}

func (s *Service) Create(ctx context.Context, req CreateRequest) (*CreateResponse, error) {
	if err := s.validateStruct(req); err != nil {
		return nil, err
	}
	var resp CreateResponse
	if err := s.doJSON(ctx, http.MethodPost, "/knowledge", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
