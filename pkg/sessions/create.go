package sessions

import (
	"context"
	"net/http"

	"github.com/gassara-kys/go-devin/pkg/types"
)

// CreateSessionRequest describes the payload used to start a session.
type CreateSessionRequest struct {
	Prompt       string   `json:"prompt" binding:"required"`
	SnapshotID   string   `json:"snapshot_id,omitempty"`
	PlaybookID   string   `json:"playbook_id,omitempty"`
	KnowledgeIDs []string `json:"knowledge_ids,omitempty"`
	SecretIDs    []string `json:"secret_ids,omitempty"`
	Tags         []string `json:"tags,omitempty"`
}

// CreateSessionResponse contains the created session metadata.
type CreateSessionResponse struct {
	types.Session
}

// Create launches a new session.
func (s *Service) Create(ctx context.Context, req CreateSessionRequest) (*CreateSessionResponse, error) {
	if err := s.validateStruct(&req); err != nil {
		return nil, err
	}
	var resp CreateSessionResponse
	if err := s.doJSON(ctx, http.MethodPost, "/sessions", nil, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
