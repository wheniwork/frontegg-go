package config

import "net/url"

func (c *FronteggConfig) GetAuthenticationServiceUrl() *url.URL {
	u := c.Endpoint.JoinPath("/auth/vendor")

	return u
}

func (c *FronteggConfig) GetIdentityServiceUrl() *url.URL {
	u := c.Endpoint.JoinPath("/identity")

	return u
}

func (c *FronteggConfig) GetAuditsServiceUrl() *url.URL {
	u := c.Endpoint.JoinPath("/audits")

	return u
}
