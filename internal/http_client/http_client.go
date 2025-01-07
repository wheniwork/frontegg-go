package http_client

import (
	"context"
	"net/http"

	"github.com/hightidecrm/frontegg-go/authenticator"
	"github.com/hightidecrm/frontegg-go/config"
)

type FronteggHttpClient struct {
	auth *authenticator.FronteggAuthenticator
	cfg  *config.FronteggConfig
}

// NewFronteggHttpClient creates a new FronteggHttpClient with the given options
func NewFronteggHttpClient(auth *authenticator.FronteggAuthenticator, cfg *config.FronteggConfig) *FronteggHttpClient {
	return &FronteggHttpClient{cfg: cfg, auth: auth}
}

func (f *FronteggHttpClient) Do(ctx context.Context, req *http.Request, tenantId *string) (*http.Response, error) {
	// get the access token
	accessToken, err := f.auth.GetAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	// set the authorization header
	req.Header.Set("x-access-token", accessToken)
	if tenantId != nil {
		req.Header.Set("frontegg-tenant-id", *tenantId)
	}

	// send the request
	return f.cfg.HttpClient.Do(req)
}
