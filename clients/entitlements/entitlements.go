package entitlements

import (
	"github.com/hightidecrm/frontegg/authenticator"
	"github.com/hightidecrm/frontegg/config"
	"github.com/hightidecrm/frontegg/internal/http_client"
)

type EntitlementsClient struct {
	cfg        *config.FronteggConfig
	auth       *authenticator.FronteggAuthenticator
	httpClient *http_client.FronteggHttpClient
}

// NewEntitlementsClient creates a new EntitlementsClient with the given options
func NewEntitlementsClient(cfg *config.FronteggConfig, auth *authenticator.FronteggAuthenticator, httpClient *http_client.FronteggHttpClient) *EntitlementsClient {
	return &EntitlementsClient{cfg: cfg, auth: auth, httpClient: httpClient}
}
