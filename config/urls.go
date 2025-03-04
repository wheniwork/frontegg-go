package config

import (
	"net/url"

	"github.com/wheniwork/frontegg-go/errors"
)

type FronteggRegion string

var (
	FronteggEU FronteggRegion = "https://api.frontegg.com"
	FronteggUS FronteggRegion = "https://api.us.frontegg.com"
	FronteggCA FronteggRegion = "https://api.ca.frontegg.com"
	FronteggAU FronteggRegion = "https://api.au.frontegg.com"
)

func (u FronteggRegion) String() string {
	return string(u)
}

func (u FronteggRegion) GetUrl() (*url.URL, error) {
	url, err := url.Parse(u.String())

	if err != nil {
		return nil, errors.ErrInvalidRegion
	}

	return url, nil
}
