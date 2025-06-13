package dto

type PolicyDTO struct {
	Identifier  string               `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Tags        []TagDTO             `json:"tags"`
	Statements  []PolicyStatementDTO `json:"statements"`
	CreatedAt   *string              `json:"created_at"`
	UpdatedAt   *string              `json:"updated_at"`
}

type PolicyStatementDTO struct {
	Effect    string   `json:"effect"`
	Actions   []string `json:"actions"`
	Resources []string `json:"resources"`
}
