package group_types

type CreateGroupRequest struct {
	Color       *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Metadata    string  `json:"metadata"`
	Name        string  `json:"name"`
}
