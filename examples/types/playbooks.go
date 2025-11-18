package types

import playbooksPkg "github.com/gassara-kys/go-devin/pkg/playbooks"

type (
	// ListPlaybooksResponse documents the list response structure for playbooks APIs.
	ListPlaybooksResponse = playbooksPkg.ListResponse
	// CreatePlaybookRequest documents the request payload for creating playbooks.
	CreatePlaybookRequest = playbooksPkg.CreateRequest
	// CreatePlaybookResponse documents the response payload for creating playbooks.
	CreatePlaybookResponse = playbooksPkg.CreateResponse
	// GetPlaybookResponse documents the response payload for fetching playbooks.
	GetPlaybookResponse = playbooksPkg.GetResponse
	// UpdatePlaybookRequest documents the request payload for updating playbooks.
	UpdatePlaybookRequest = playbooksPkg.UpdateRequest
	// UpdatePlaybookResponse documents the response payload for updating playbooks.
	UpdatePlaybookResponse = playbooksPkg.UpdateResponse
	// DeletePlaybookResponse documents the response payload for deleting playbooks.
	DeletePlaybookResponse = playbooksPkg.DeleteResponse
)
