package sessions

import (
	"context"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/gassara-kys/go-devin/pkg/types"
)

type ListSessionsRequest struct {
	Limit  int      `url:"limit,omitempty" binding:"omitempty,min=1,max=1000"`
	Offset int      `url:"offset,omitempty" binding:"omitempty,min=0"`
	Tags   []string `url:"tags,omitempty" binding:"omitempty,dive,required"`
}

type ListSessionsResponse struct {
	Sessions []types.Session `json:"sessions"`
}

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
