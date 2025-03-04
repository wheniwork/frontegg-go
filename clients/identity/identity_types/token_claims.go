package identity_types

import "github.com/golang-jwt/jwt/v5"

type FronteggBaseTokenClaims struct {
	jwt.RegisteredClaims

	Type AccessTokenType `json:"type"`

	Metadata map[string]interface{} `json:"metadata"`

	TenantId      string   `json:"tenantId"`
	ApplicationID string   `json:"applicationId"`
	Permissions   []string `json:"permissions"`
	Roles         []string `json:"roles"`
}

func (f *FronteggBaseTokenClaims) GetTokenType() AccessTokenType {
	return f.Type
}

func (f *FronteggBaseTokenClaims) IsTenantToken() (bool, *FronteggTenantTokenClaims) {
	return false, nil
}

func (f *FronteggBaseTokenClaims) IsUserToken() (bool, *FronteggUserTokenClaims) {
	return false, nil
}

type FronteggTenantTokenClaims struct {
	FronteggBaseTokenClaims
	CreatedByUserID string `json:"createdByUserId"`
}

func (f *FronteggTenantTokenClaims) GetTokenType() AccessTokenType {
	return f.Type
}

func (f *FronteggTenantTokenClaims) IsTenantToken() (bool, *FronteggTenantTokenClaims) {
	return true, f
}

func (f *FronteggTenantTokenClaims) IsUserToken() (bool, *FronteggUserTokenClaims) {
	return false, nil
}

type FronteggUserTokenClaims struct {
	FronteggBaseTokenClaims
	Name          string `json:"name"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"emailVerified"`

	TenantIds []string `json:"tenantIds"`

	ProfilePictureUrl string `json:"profilePictureUrl"`
}

func (f *FronteggUserTokenClaims) GetTokenType() AccessTokenType {
	return f.Type
}

func (f *FronteggUserTokenClaims) IsTenantToken() (bool, *FronteggTenantTokenClaims) {
	return false, nil
}

func (f *FronteggUserTokenClaims) IsUserToken() (bool, *FronteggUserTokenClaims) {
	return true, f
}
