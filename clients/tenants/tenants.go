package tenants

import (
	"github.com/wheniwork/frontegg-go/authenticator"
	"github.com/wheniwork/frontegg-go/config"
	"github.com/wheniwork/frontegg-go/internal/http_client"
)

var (
	baseEndpoint   = "/tenants/resources/tenants/v2"
	v1BaseEndpoint = "/tenants/resources/tenants/v1"
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
