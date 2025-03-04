package groups

import (
	"github.com/wheniwork/frontegg-go/authenticator"
	"github.com/wheniwork/frontegg-go/config"
	"github.com/wheniwork/frontegg-go/internal/http_client"
)

var (
	baseEndpoint   = "/identity/resources/groups/v1"
	v2BaseEndpoint = "/identity/resources/groups/v2"
)

type GroupClient struct {
	cfg        *config.FronteggConfig
	auth       *authenticator.FronteggAuthenticator
	httpClient *http_client.FronteggHttpClient
}

// NewGroupClient creates a new GroupClient with the given options
func NewGroupClient(cfg *config.FronteggConfig, auth *authenticator.FronteggAuthenticator, httpClient *http_client.FronteggHttpClient) *GroupClient {
	return &GroupClient{cfg: cfg, auth: auth, httpClient: httpClient}
}
