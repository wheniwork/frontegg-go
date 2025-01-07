package identity

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"net/http"

	"github.com/hightidecrm/frontegg/authenticator"
	"github.com/hightidecrm/frontegg/clients/identity/identity_types"
	"github.com/hightidecrm/frontegg/config"
	"github.com/hightidecrm/frontegg/errors"
	"github.com/hightidecrm/frontegg/internal/http_client"
)

type IdentityClient struct {
	cfg        *config.FronteggConfig
	auth       *authenticator.FronteggAuthenticator
	httpClient *http_client.FronteggHttpClient

	publicKey *rsa.PublicKey
}

// NewIdentityClient creates a new IdentityClient with the given options
func NewIdentityClient(cfg *config.FronteggConfig, auth *authenticator.FronteggAuthenticator, httpClient *http_client.FronteggHttpClient) *IdentityClient {
	return &IdentityClient{cfg: cfg, auth: auth, httpClient: httpClient}
}

func (i *IdentityClient) GetPublicKey(ctx context.Context, ignoreCached bool) (*rsa.PublicKey, error) {
	if !ignoreCached && i.publicKey != nil {
		return i.publicKey, nil
	}

	fullUrl := i.cfg.GetIdentityServiceUrl()

	fullUrl = fullUrl.JoinPath("/resources/configurations/v1")

	req, err := http.NewRequestWithContext(ctx, "GET", fullUrl.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := i.httpClient.Do(ctx, req, nil)

	if err != nil {
		return nil, err
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.ErrInvalidResponse
	}

	// decode the response
	var value identity_types.IdentityConfiguration

	err = json.NewDecoder(resp.Body).Decode(&value)

	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(value.PublicKey))
	if block == nil {
		return nil, errors.ErrAuthenticationFailed
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.ErrAuthenticationFailed
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.ErrAuthenticationFailed
	}

	i.publicKey = rsaPub

	return rsaPub, nil
}
