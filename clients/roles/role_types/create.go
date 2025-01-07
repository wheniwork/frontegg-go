package role_types

type CreateRoleRequest struct {
	Key           string   `json:"key"`
	Name          string   `json:"name"`
	Description   *string  `json:"description,omitempty"`
	IsDefault     *bool    `json:"isDefault,omitempty"`
	BaseRoleID    string   `json:"baseRoleId"`
	PermissionIds []string `json:"permissionIds"`
}
