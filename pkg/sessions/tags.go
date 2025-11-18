package sessions

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// UpdateTagsRequest contains the desired set of tags for a session.
type UpdateTagsRequest struct {
	Tags []string `json:"tags" binding:"required,min=1,dive,required"`
}

// UpdateTagsResponse contains the API's confirmation message.
type UpdateTagsResponse struct {
	Detail string `json:"detail"`
}

// UpdateTags replaces the tags for the specified session.
func (s *Service) UpdateTags(ctx context.Context, sessionID string, req UpdateTagsRequest) (*UpdateTagsResponse, error) {
	if sessionID == "" {
		return nil, fmt.Errorf("sessionID is required")
	}
	if err := s.validateStruct(&req); err != nil {
		return nil, err
	}
	var resp UpdateTagsResponse
	path := fmt.Sprintf("/sessions/%s/tags", url.PathEscape(sessionID))
	if err := s.doJSON(ctx, http.MethodPut, path, nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
