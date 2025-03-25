package identity_types

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"slices"
)

type ActorClaims struct {
	Subject string `json:"sub"`
	Type    string `json:"type"`
}

type FronteggBaseTokenClaims struct {
	jwt.RegisteredClaims

	Type AccessTokenType `json:"type"`

	Metadata map[string]interface{} `json:"metadata"`

	TenantId      string   `json:"tenantId"`
	ApplicationID string   `json:"applicationId"`
	Permissions   []string `json:"permissions"`
	Roles         []string `json:"roles"`
}

func (f *FronteggBaseTokenClaims) Validate() error {
	if f.Type == "" {
		return errors.New("type is required")
	}

	if f.TenantId == "" {
		return errors.New("tenantId is required")
	}

	return nil
}

func (f *FronteggBaseTokenClaims) GetTokenType() AccessTokenType {
	return f.Type
}

// HasPermission checks if the token has a specific permission.
// It supports exact matches and wildcard permissions (ending with '*').
// For example, if the token has permission "fe.secure.*", it will match "fe.secure.read.users".
func (f *FronteggBaseTokenClaims) HasPermission(permission string) bool {
	for _, p := range f.Permissions {
		// Check for exact match
		if p == permission {
			return true
		}

		// Check for wildcard match
		if len(p) > 0 && p[len(p)-1] == '*' {
			prefix := p[:len(p)-1]
			if len(permission) >= len(prefix) && permission[:len(prefix)] == prefix {
				return true
			}
		}
	}
	return false
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

func (f *FronteggTenantTokenClaims) Validate() error {
	return f.FronteggBaseTokenClaims.Validate()
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

	ProfilePictureUrl string      `json:"profilePictureUrl"`
	Actor             ActorClaims `json:"act"`
	Acr               string      `json:"acr,omitempty"`
	Amr               []string    `json:"amr,omitempty"`
	AuthTime          int64       `json:"authTime,omitempty"`
}

func (f *FronteggUserTokenClaims) Validate() error {
	return f.FronteggBaseTokenClaims.Validate()
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

func (f *FronteggUserTokenClaims) IsSteppedUp() (bool, *FronteggUserTokenClaims) {
	isACRValid := f.Acr == StepUpAcrValue
	isAMRValid := f.Amr != nil && len(f.Amr) > 0 && slices.Contains(f.Amr, StepUpMFAValue)

	return isACRValid && isAMRValid, f
}
