package identity_types

type FronteggTokenInterface interface {
	GetTokenType() AccessTokenType
	IsTenantToken() (bool, *FronteggTenantTokenClaims)
	IsUserToken() (bool, *FronteggUserTokenClaims)
}
