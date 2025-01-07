package users

import (
	"github.com/hightidecrm/frontegg-go/authenticator"
	"github.com/hightidecrm/frontegg-go/config"
	"github.com/hightidecrm/frontegg-go/internal/http_client"
)

var (
	vendorOnlyEndpoint = "/identity/resources/vendor-only/users/v1"
	tenantEndpoint     = "/identity/resources/users/v1/"
)

type UserClient struct {
	cfg        *config.FronteggConfig
	auth       *authenticator.FronteggAuthenticator
	httpClient *http_client.FronteggHttpClient
}

// NewUserClient creates a new UserClient with the given options
func NewUserClient(cfg *config.FronteggConfig, auth *authenticator.FronteggAuthenticator, httpClient *http_client.FronteggHttpClient) *UserClient {
	return &UserClient{cfg: cfg, auth: auth, httpClient: httpClient}
}
