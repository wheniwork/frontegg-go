package group_types

type RolesRequest struct {
	RoleIds []string `json:"roleIds"`
}

type UsersRequest struct {
	UserIds []string `json:"userIds"`
}
