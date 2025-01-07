package users

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hightidecrm/frontegg/clients/users/user_types"
	"github.com/hightidecrm/frontegg/errors"
	"github.com/hightidecrm/frontegg/internal/http_client"
)

func (u *UserClient) GetUser(ctx context.Context, userId string) (*user_types.User, error) {
	fullUrl := u.cfg.Endpoint.JoinPath(vendorOnlyEndpoint, userId)

	req, err := http.NewRequestWithContext(ctx, "GET", fullUrl.String(), nil)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to get user", err)
	}

	var user user_types.User

	resp, err := u.httpClient.Do(ctx, req, nil)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to get user", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to get user", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&user)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode get user response body", err)
	}

	return &user, nil
}

func (u *UserClient) GetTenantUser(ctx context.Context, tenantId string, userId string) (*user_types.User, error) {
	fullUrl := u.cfg.Endpoint.JoinPath(tenantEndpoint, userId)

	req, err := http.NewRequestWithContext(ctx, "GET", fullUrl.String(), nil)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to get tenant user", err)
	}

	var user user_types.User

	resp, err := u.httpClient.Do(ctx, req, &tenantId)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to get tenant user", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to get tenant user", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&user)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode get tenant user response body", err)
	}

	return &user, nil
}

func (u *UserClient) GetUserByEmail(ctx context.Context, tenantId string, email string) (*user_types.User, error) {
	fullUrl := u.cfg.Endpoint.JoinPath(tenantEndpoint, "/email")

	q := fullUrl.Query()

	q.Add("email", email)

	fullUrl.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", fullUrl.String(), nil)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to get tenant user by email", err)
	}

	var user user_types.User

	resp, err := u.httpClient.Do(ctx, req, &tenantId)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to get tenant user by email", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to get tenant user by email", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&user)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode get tenant user response body", err)
	}

	return &user, nil
}
