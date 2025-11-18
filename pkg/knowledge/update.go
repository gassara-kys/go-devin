package knowledge

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gassara-kys/go-devin/pkg/types"
)

type UpdateRequest struct {
	Name               *string `json:"name,omitempty"`
	Body               *string `json:"body,omitempty"`
	TriggerDescription *string `json:"trigger_description,omitempty"`
	ParentFolderID     *string `json:"parent_folder_id,omitempty"`
	PinnedRepo         *string `json:"pinned_repo,omitempty"`
}

type UpdateResponse struct {
	types.Knowledge
}

func (s *Service) Update(ctx context.Context, noteID string, req UpdateRequest) (*UpdateResponse, error) {
	if noteID == "" {
		return nil, fmt.Errorf("noteID is required")
	}
	if req.Name == nil && req.Body == nil && req.TriggerDescription == nil && req.ParentFolderID == nil && req.PinnedRepo == nil {
		return nil, fmt.Errorf("at least one field must be provided")
	}
	if err := s.validateStruct(req); err != nil {
		return nil, err
	}
	var resp UpdateResponse
	path := fmt.Sprintf("/knowledge/%s", url.PathEscape(noteID))
	if err := s.doJSON(ctx, http.MethodPut, path, nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
