package sessions

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type SendMessageRequest struct {
	Message       string   `json:"message" binding:"required"`
	AttachmentIDs []string `json:"attachment_ids,omitempty"`
}

type SendMessageResponse struct {
	MessageID string `json:"message_id"`
	Status    string `json:"status"`
	Detail    string `json:"detail"`
}

func (s *Service) SendMessage(ctx context.Context, sessionID string, req SendMessageRequest) (*SendMessageResponse, error) {
	if sessionID == "" {
		return nil, fmt.Errorf("sessionID is required")
	}
	if err := s.validateStruct(&req); err != nil {
		return nil, err
	}
	var resp SendMessageResponse
	path := fmt.Sprintf("/sessions/%s/messages", url.PathEscape(sessionID))
	if err := s.doJSON(ctx, http.MethodPost, path, nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
