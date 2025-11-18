package types

// Secret represents a stored credential in Devin.
type Secret struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Key       string `json:"key"`
	CreatedAt string `json:"created_at"`
}
