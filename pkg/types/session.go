package types

import "encoding/json"

// Session represents a high-level Devin session record.
type Session struct {
	SessionID        string          `json:"session_id"`
	Status           string          `json:"status"`
	Title            string          `json:"title"`
	CreatedAt        string          `json:"created_at"`
	UpdatedAt        string          `json:"updated_at"`
	SnapshotID       *string         `json:"snapshot_id,omitempty"`
	PlaybookID       *string         `json:"playbook_id,omitempty"`
	PullRequest      *PullRequest    `json:"pull_request,omitempty"`
	StructuredOutput json.RawMessage `json:"structured_output,omitempty"`
	StatusEnum       string          `json:"status_enum,omitempty"`
	Tags             []string        `json:"tags,omitempty"`
}

// PullRequest mirrors the metadata returned for linked pull requests.
type PullRequest struct {
	URL string `json:"url"`
}

// SessionMessage describes a single message inside a session transcript.
type SessionMessage struct {
	MessageID   string              `json:"message_id"`
	Sender      string              `json:"sender"`
	Content     string              `json:"content"`
	Status      string              `json:"status"`
	CreatedAt   string              `json:"created_at"`
	Attachments []SessionAttachment `json:"attachments,omitempty"`
}

// SessionAttachment stores metadata for files attached to a message.
type SessionAttachment struct {
	AttachmentID string `json:"attachment_id"`
	FileName     string `json:"file_name"`
}

// SessionDetail combines the base session info with optional messages.
type SessionDetail struct {
	Session
	Messages []SessionMessage `json:"messages,omitempty"`
}
