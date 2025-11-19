package sessions

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// SendMessageRequest contains the payload for sending a session message.
type SendMessageRequest struct {
	Message       string   `json:"message" binding:"required"`
	AttachmentIDs []string `json:"attachment_ids,omitempty"`
}

// SendMessage posts a new message to a session.
func (s *Service) SendMessage(ctx context.Context, sessionID string, req SendMessageRequest) error {
	if sessionID == "" {
		return fmt.Errorf("sessionID is required")
	}
	if err := s.validateStruct(&req); err != nil {
		return err
	}

	path := fmt.Sprintf("/sessions/%s/message", url.PathEscape(sessionID))
	if err := s.doJSON(ctx, http.MethodPost, path, nil, req, nil); err != nil {
		return err
	}
	return nil
}
