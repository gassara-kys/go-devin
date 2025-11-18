package playbooks

import (
	"context"
	"net/http"

	"github.com/gassara-kys/go-devin/pkg/types"
)

// ListResponse contains the results of listing playbooks.
type ListResponse struct {
	Playbooks []types.Playbook `json:"playbooks"`
}

// List fetches all playbooks accessible to the caller.
func (s *Service) List(ctx context.Context) (*ListResponse, error) {
	var records []types.Playbook
	if err := s.doJSON(ctx, http.MethodGet, "/playbooks", nil, nil, &records); err != nil {
		return nil, err
	}
	return &ListResponse{Playbooks: records}, nil
}
