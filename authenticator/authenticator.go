package authenticator

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/wheniwork/frontegg-go/config"
	"github.com/wheniwork/frontegg-go/errors"
)

type FronteggAuthenticator struct {
	cfg *config.FronteggConfig

	accessToken       string
	accessTokenExpiry *time.Time
}

func NewFronteggAuthenticator(cfg *config.FronteggConfig) *FronteggAuthenticator {
	return &FronteggAuthenticator{cfg: cfg}
}

// Authenticate authenticates the client with Frontegg
func (a *FronteggAuthenticator) Authenticate(ctx context.Context) error {
	// authenticate the client
	url := a.cfg.GetAuthenticationServiceUrl()

	// get the access token
	body := AuthenticateRequest{
		ClientID: a.cfg.ClientID,
		Secret:   a.cfg.ApiKey,
	}

	// send the request
	rawData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// send the request
	req, err := http.NewRequestWithContext(ctx, "POST", url.String(), bytes.NewBuffer(rawData))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	// send the request
	resp, err := a.cfg.HttpClient.Do(req)
	if err != nil {
		return err
	}

	// handle the response
	if resp.StatusCode != 200 {
		return errors.ErrAuthenticationFailed
	}

	// decode the response
	defer resp.Body.Close()
	var authResponse AuthenticateResponse

	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return err
	}

	// set the access token and expiry
	a.accessToken = authResponse.Token
	expiry := time.Now().Add(time.Duration(authResponse.ExpiresIn) * 1000 * 90 / 100)

	a.accessTokenExpiry = &expiry

	return nil
}

// GetAccessToken returns the access token
func (a *FronteggAuthenticator) GetAccessToken(ctx context.Context) (string, error) {
	if a.accessTokenExpiry == nil || a.accessTokenExpiry.Before(time.Now()) {
		if err := a.Authenticate(ctx); err != nil {
			return "", err
		}
	}

	return a.accessToken, nil
}
