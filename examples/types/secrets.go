package types

import secretsPkg "github.com/gassara-kys/go-devin/pkg/secrets"

type (
	// ListSecretsResponse documents the list response structure for secrets APIs.
	ListSecretsResponse = secretsPkg.ListResponse
	// CreateSecretRequest documents the request payload for creating secrets.
	CreateSecretRequest = secretsPkg.CreateRequest
	// CreateSecretResponse documents the response payload for creating secrets.
	CreateSecretResponse = secretsPkg.CreateResponse
)
