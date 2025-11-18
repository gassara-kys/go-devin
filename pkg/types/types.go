package types

import "encoding/json"

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

type PullRequest struct {
	URL string `json:"url"`
}

type SessionMessage struct {
	MessageID   string              `json:"message_id"`
	Sender      string              `json:"sender"`
	Content     string              `json:"content"`
	Status      string              `json:"status"`
	CreatedAt   string              `json:"created_at"`
	Attachments []SessionAttachment `json:"attachments,omitempty"`
}

type SessionAttachment struct {
	AttachmentID string `json:"attachment_id"`
	FileName     string `json:"file_name"`
}

type SessionDetail struct {
	Session
	Messages []SessionMessage `json:"messages,omitempty"`
}

type Secret struct {
	SecretID   string `json:"secret_id"`
	SecretType string `json:"secret_type"`
	SecretName string `json:"secret_name"`
	CreatedAt  string `json:"created_at"`
}

type Knowledge struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Body               string `json:"body"`
	TriggerDescription string `json:"trigger_description"`
	ParentFolderID     string `json:"parent_folder_id"`
	CreatedAt          string `json:"created_at"`
	PinnedRepo         string `json:"pinned_repo,omitempty"`
}

type KnowledgeFolder struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type Playbook struct {
	PlaybookID        string `json:"playbook_id"`
	Title             string `json:"title"`
	Body              string `json:"body"`
	Status            string `json:"status"`
	AccessType        string `json:"access_type"`
	OrgID             string `json:"org_id"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	CreatedByUserID   string `json:"created_by_user_id"`
	CreatedByUserName string `json:"created_by_user_name"`
	UpdatedByUserID   string `json:"updated_by_user_id"`
	UpdatedByUserName string `json:"updated_by_user_name"`
	Macro             string `json:"macro"`
}
