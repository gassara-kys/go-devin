package playbooks

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// UpdateRequest carries the fields that can be modified for a playbook.
type UpdateRequest struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	Macro string `json:"macro,omitempty"`
}

// UpdateResponse reports the resulting status of the update operation.
type UpdateResponse struct {
	Status string `json:"status"`
}

// Update modifies a playbook using the provided fields.
func (s *Service) Update(ctx context.Context, playbookID string, req UpdateRequest) (*UpdateResponse, error) {
	if playbookID == "" {
		return nil, fmt.Errorf("playbookID is required")
	}
	if req.Title == "" && req.Body == "" && req.Macro == "" {
		return nil, fmt.Errorf("at least one field must be provided")
	}
	if err := s.validateStruct(req); err != nil {
		return nil, err
	}
	var resp UpdateResponse
	path := fmt.Sprintf("/playbooks/%s", url.PathEscape(playbookID))
	if err := s.doJSON(ctx, http.MethodPut, path, nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
