package types

// Playbook models the automation scripts exposed by Devin.
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
