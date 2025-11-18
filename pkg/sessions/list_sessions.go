package sessions

import (
	"context"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/gassara-kys/go-devin/pkg/types"
)

// ListSessionsRequest contains filters for listing sessions.
type ListSessionsRequest struct {
	Limit  int      `url:"limit,omitempty" binding:"omitempty,min=1,max=1000"`
	Offset int      `url:"offset,omitempty" binding:"omitempty,min=0"`
	Tags   []string `url:"tags,omitempty" binding:"omitempty,dive,required"`
}

// ListSessionsResponse contains the paginated sessions list.
type ListSessionsResponse struct {
	Sessions []types.Session `json:"sessions"`
}

// List retrieves sessions optionally filtered by tags.
func (s *Service) List(ctx context.Context, req *ListSessionsRequest) (*ListSessionsResponse, error) {
	var (
		q   url.Values
		err error
	)
	if req != nil {
		if err := s.validateStruct(req); err != nil {
			return nil, err
		}
		q, err = query.Values(req)
		if err != nil {
			return nil, err
		}
	}
	var resp ListSessionsResponse
	if err := s.doJSON(ctx, http.MethodGet, "/sessions", q, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
