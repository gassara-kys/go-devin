package playbooks

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type DeleteResponse struct {
	Status string `json:"status"`
}

func (s *Service) Delete(ctx context.Context, playbookID string) (*DeleteResponse, error) {
	if playbookID == "" {
		return nil, fmt.Errorf("playbookID is required")
	}
	var resp DeleteResponse
	path := fmt.Sprintf("/playbooks/%s", url.PathEscape(playbookID))
	if err := s.doJSON(ctx, http.MethodDelete, path, nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
