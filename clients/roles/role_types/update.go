package role_types

type UpdateRoleRequest struct {
	IsDefault     *bool   `json:"isDefault,omitempty"`
	FirstUserRole *bool   `json:"firstUserRole,omitempty"`
	MigrateRole   *bool   `json:"migrateRole,omitempty"`
	Level         *int    `json:"level,omitempty"`
	Key           *string `json:"key,omitempty"`
	Name          *string `json:"name,omitempty"`
	Description   *string `json:"description,omitempty"`
}

type AssignPermissionsToRoleRequest struct {
	PermissionIds []string `json:"permissionIds"`
}
