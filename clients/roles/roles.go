package roles

import (
	"github.com/hightidecrm/frontegg/authenticator"
	"github.com/hightidecrm/frontegg/config"
	"github.com/hightidecrm/frontegg/internal/http_client"
)

var (
	baseEndpoint   = "/identity/resources/roles/v2"
	v1BaseEndpoint = "/identity/resources/roles/v1"
)

type RoleClient struct {
	cfg        *config.FronteggConfig
	auth       *authenticator.FronteggAuthenticator
	httpClient *http_client.FronteggHttpClient
}

// NewRoleClient creates a new RoleClient with the given options
func NewRoleClient(cfg *config.FronteggConfig, auth *authenticator.FronteggAuthenticator, httpClient *http_client.FronteggHttpClient) *RoleClient {
	return &RoleClient{cfg: cfg, auth: auth, httpClient: httpClient}
}
