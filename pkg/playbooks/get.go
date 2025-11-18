package playbooks

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gassara-kys/go-devin/pkg/types"
)

// GetResponse contains a single playbook record.
type GetResponse struct {
	types.Playbook
}

// Get retrieves a playbook by ID.
func (s *Service) Get(ctx context.Context, playbookID string) (*GetResponse, error) {
	if playbookID == "" {
		return nil, fmt.Errorf("playbookID is required")
	}
	var resp GetResponse
	path := fmt.Sprintf("/playbooks/%s", url.PathEscape(playbookID))
	if err := s.doJSON(ctx, http.MethodGet, path, nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
