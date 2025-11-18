package types

import secretsPkg "github.com/gassara-kys/go-devin/pkg/secrets"

type (
	ListSecretsResponse  = secretsPkg.ListResponse
	CreateSecretRequest  = secretsPkg.CreateRequest
	CreateSecretResponse = secretsPkg.CreateResponse
)
