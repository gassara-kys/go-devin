package knowledge

import (
	"context"
	"net/http"

	"github.com/gassara-kys/go-devin/pkg/types"
)

// ListResponse captures the data returned from the list endpoint.
type ListResponse struct {
	Knowledge []types.Knowledge       `json:"knowledge"`
	Folders   []types.KnowledgeFolder `json:"folders"`
}

// List fetches all knowledge entries and folders.
func (s *Service) List(ctx context.Context) (*ListResponse, error) {
	var resp ListResponse
	if err := s.doJSON(ctx, http.MethodGet, "/knowledge", nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
