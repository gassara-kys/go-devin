package types

import playbooksPkg "github.com/gassara-kys/go-devin/pkg/playbooks"

type (
	ListPlaybooksResponse  = playbooksPkg.ListResponse
	CreatePlaybookRequest  = playbooksPkg.CreateRequest
	CreatePlaybookResponse = playbooksPkg.CreateResponse
	GetPlaybookResponse    = playbooksPkg.GetResponse
	UpdatePlaybookRequest  = playbooksPkg.UpdateRequest
	UpdatePlaybookResponse = playbooksPkg.UpdateResponse
	DeletePlaybookResponse = playbooksPkg.DeleteResponse
)
