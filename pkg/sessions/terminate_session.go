package sessions

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type TerminateSessionResponse struct {
	Detail string `json:"detail"`
}

func (s *Service) Terminate(ctx context.Context, sessionID string) (*TerminateSessionResponse, error) {
	if sessionID == "" {
		return nil, fmt.Errorf("sessionID is required")
	}
	var resp TerminateSessionResponse
	path := fmt.Sprintf("/sessions/%s", url.PathEscape(sessionID))
	if err := s.doJSON(ctx, http.MethodDelete, path, nil, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
