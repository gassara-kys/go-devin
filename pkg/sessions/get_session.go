package sessions

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gassara-kys/go-devin/pkg/types"
)

func (s *Service) Get(ctx context.Context, sessionID string) (*types.SessionDetail, error) {
	if sessionID == "" {
		return nil, fmt.Errorf("sessionID is required")
	}
	var resp types.SessionDetail
	path := fmt.Sprintf("/sessions/%s", url.PathEscape(sessionID))
	if err := s.doJSON(ctx, http.MethodGet, path, nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
