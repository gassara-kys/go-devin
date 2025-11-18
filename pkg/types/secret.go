package types

// Secret represents a stored credential in Devin.
type Secret struct {
	SecretID   string `json:"secret_id"`
	SecretType string `json:"secret_type"`
	SecretName string `json:"secret_name"`
	CreatedAt  string `json:"created_at"`
}
