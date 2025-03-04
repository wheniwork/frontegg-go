package events

import (
	"github.com/wheniwork/frontegg-go/authenticator"
	"github.com/wheniwork/frontegg-go/config"
	"github.com/wheniwork/frontegg-go/internal/http_client"
)

type EventsClient struct {
	cfg        *config.FronteggConfig
	auth       *authenticator.FronteggAuthenticator
	httpClient *http_client.FronteggHttpClient
}

// NewEventsClient creates a new EventsClient with the given options
func NewEventsClient(cfg *config.FronteggConfig, auth *authenticator.FronteggAuthenticator, httpClient *http_client.FronteggHttpClient) *EventsClient {
	return &EventsClient{cfg: cfg, auth: auth, httpClient: httpClient}
}
