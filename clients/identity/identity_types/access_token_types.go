package identity_types

type AccessTokenType string

const (
	AccessTokenTypeUser    AccessTokenType = "user"
	AccessTokenTypeTenant  AccessTokenType = "tenantApiToken"
	AccessTokenTypeUserApi AccessTokenType = "userApiToken"
)
