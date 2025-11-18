package types

// Knowledge represents an individual knowledge entry.
type Knowledge struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Body               string `json:"body"`
	TriggerDescription string `json:"trigger_description"`
	ParentFolderID     string `json:"parent_folder_id"`
	CreatedAt          string `json:"created_at"`
	PinnedRepo         string `json:"pinned_repo,omitempty"`
}

// KnowledgeFolder captures hierarchy metadata for knowledge entries.
type KnowledgeFolder struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
