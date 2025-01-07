package config

import "net/url"

type FronteggOption func(*FronteggConfig) error

// WithRegion sets the region of the FronteggConfig
func WithRegion(region FronteggRegion) FronteggOption {
	return func(c *FronteggConfig) error {
		c.Region = region
		var err error
		c.Endpoint, err = region.GetUrl()
		return err
	}
}

// WithClientID sets the clientID of the FronteggConfig
func WithClientID(clientID string) FronteggOption {
	return func(c *FronteggConfig) error {
		c.ClientID = clientID
		return nil
	}
}

// WithApiKey sets the apiKey of the FronteggConfig
func WithApiKey(apiKey string) FronteggOption {
	return func(c *FronteggConfig) error {
		c.ApiKey = apiKey
		return nil
	}
}

// WithEndpoint sets the endpoint of the FronteggConfig
func WithEndpoint(endpoint *url.URL) FronteggOption {
	return func(c *FronteggConfig) error {
		c.Endpoint = endpoint
		return nil
	}
}
