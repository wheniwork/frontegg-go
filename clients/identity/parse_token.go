package identity

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wheniwork/frontegg-go/clients/identity/identity_types"
	"github.com/wheniwork/frontegg-go/errors"
)

func (i *IdentityClient) getUserClaims(ctx context.Context, token string) (*identity_types.FronteggUserTokenClaims, error) {
	var claims identity_types.FronteggUserTokenClaims

	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return i.GetPublicKey(ctx, false)
	})

	if err != nil {
		return nil, err
	}

	return &claims, nil
}

func (i *IdentityClient) getTenantClaims(ctx context.Context, token string) (*identity_types.FronteggTenantTokenClaims, error) {
	var claims identity_types.FronteggTenantTokenClaims

	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return i.GetPublicKey(ctx, false)
	})

	if err != nil {
		return nil, err
	}

	return &claims, nil
}

func (i *IdentityClient) ParseToken(ctx context.Context, token string) (identity_types.FronteggTokenInterface, error) {
	// parse the token
	var claims identity_types.FronteggBaseTokenClaims

	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return i.GetPublicKey(ctx, false)
	})

	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.ErrInvalidToken
	}

	var switchedClaims identity_types.FronteggTokenInterface

	switch claims.GetTokenType() {
	case identity_types.AccessTokenTypeTenant:
		switchedClaims, err = i.getTenantClaims(ctx, token)
		break
	case identity_types.AccessTokenTypeUserApi:
		fallthrough
	case identity_types.AccessTokenTypeUser:
		switchedClaims, err = i.getUserClaims(ctx, token)
		break
	default:
		switchedClaims = &claims
	}

	if err != nil {
		return nil, err
	}

	return switchedClaims, nil
}
