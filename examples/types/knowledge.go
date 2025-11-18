package types

import knowledgePkg "github.com/gassara-kys/go-devin/pkg/knowledge"

type (
	// ListKnowledgeResponse documents the list response shape for knowledge APIs.
	ListKnowledgeResponse = knowledgePkg.ListResponse
	// CreateKnowledgeRequest documents the create request payload for knowledge APIs.
	CreateKnowledgeRequest = knowledgePkg.CreateRequest
	// CreateKnowledgeResponse documents the create response payload for knowledge APIs.
	CreateKnowledgeResponse = knowledgePkg.CreateResponse
	// UpdateKnowledgeRequest documents the update request payload for knowledge APIs.
	UpdateKnowledgeRequest = knowledgePkg.UpdateRequest
	// UpdateKnowledgeResponse documents the update response payload for knowledge APIs.
	UpdateKnowledgeResponse = knowledgePkg.UpdateResponse
)
