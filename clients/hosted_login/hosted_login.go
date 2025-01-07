package hosted_login

import (
	"github.com/hightidecrm/frontegg-go/authenticator"
	"github.com/hightidecrm/frontegg-go/config"
	"github.com/hightidecrm/frontegg-go/internal/http_client"
)

type HostedLoginClient struct {
	cfg        *config.FronteggConfig
	auth       *authenticator.FronteggAuthenticator
	httpClient *http_client.FronteggHttpClient
}

// NewHostedLoginClient creates a new HostedLoginClient with the given options
func NewHostedLoginClient(cfg *config.FronteggConfig, auth *authenticator.FronteggAuthenticator, httpClient *http_client.FronteggHttpClient) *HostedLoginClient {
	return &HostedLoginClient{cfg: cfg, auth: auth, httpClient: httpClient}
}
