package config

import (
	"net/http"
	"net/url"
	"os"
)

// FronteggConfig holds the configuration for the Frontegg client
type FronteggConfig struct {
	Region FronteggRegion

	Endpoint *url.URL

	ClientID string

	ApiKey string

	HttpClient *http.Client
}

// NewFronteggConfig creates a new FronteggConfig with the given options
// If no options are provided, the default values are used
// The default values are:
// - Region: FronteggUS
// - Endpoint: https://api.us.frontegg.com
// - ApiKey: the value of the FRONTEGG_API_KEY environment variable
// - ClientID: the value of the FRONTEGG_CLIENT_ID environment variable
func NewFronteggConfig(opts ...FronteggOption) (*FronteggConfig, error) {
	defaultEndpoint, err := FronteggUS.GetUrl()
	if err != nil {
		return nil, err
	}

	cfg := &FronteggConfig{
		Region:     FronteggUS,
		Endpoint:   defaultEndpoint,
		ApiKey:     os.Getenv("FRONTEGG_API_KEY"),
		ClientID:   os.Getenv("FRONTEGG_CLIENT_ID"),
		HttpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}
