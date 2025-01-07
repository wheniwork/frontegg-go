package tenants

import (
	"github.com/hightidecrm/frontegg/authenticator"
	"github.com/hightidecrm/frontegg/config"
	"github.com/hightidecrm/frontegg/internal/http_client"
)

var (
	baseEndpoint = "/tenants/resources/tenants/v2"
)

type TenantClient struct {
	cfg        *config.FronteggConfig
	auth       *authenticator.FronteggAuthenticator
	httpClient *http_client.FronteggHttpClient
}

// NewTenantClient creates a new TenantClient with the given options
func NewTenantClient(cfg *config.FronteggConfig, auth *authenticator.FronteggAuthenticator, httpClient *http_client.FronteggHttpClient) *TenantClient {
	return &TenantClient{cfg: cfg, auth: auth, httpClient: httpClient}
}
