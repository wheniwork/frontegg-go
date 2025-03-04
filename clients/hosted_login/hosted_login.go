package hosted_login

import (
	"github.com/wheniwork/frontegg-go/authenticator"
	"github.com/wheniwork/frontegg-go/config"
	"github.com/wheniwork/frontegg-go/internal/http_client"
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
