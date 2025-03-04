package users

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/wheniwork/frontegg-go/clients/users/user_types"
	"github.com/wheniwork/frontegg-go/errors"
	"github.com/wheniwork/frontegg-go/internal/http_client"
)

func (u *UserClient) CreateUser(ctx context.Context, tenantId string, body *user_types.CreateUserRequest) (*user_types.User, error) {
	fullUrl := u.cfg.Endpoint.JoinPath(tenantEndpoint)

	parsedBody, err := json.Marshal(body)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to marshal create user request body", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fullUrl.String(), bytes.NewBuffer(parsedBody))

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to create user", err)
	}

	req.Header.Set("Content-Type", "application/json")

	var user user_types.User

	resp, err := u.httpClient.Do(ctx, req, &tenantId)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to create user", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to create user", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&user)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode create user response body", err)
	}

	return &user, nil
}

func (u *UserClient) UpdateUser(ctx context.Context, userId string, body *user_types.UpdateUserRequest) (*user_types.User, error) {
	fullUrl := u.cfg.Endpoint.JoinPath(vendorOnlyEndpoint, userId)

	parsedBody, err := json.Marshal(body)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to marshal update user request body", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", fullUrl.String(), bytes.NewBuffer(parsedBody))

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to update user", err)
	}

	req.Header.Set("Content-Type", "application/json")

	var user user_types.User

	resp, err := u.httpClient.Do(ctx, req, nil)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to update user", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to update user", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&user)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode update user response body", err)
	}

	return &user, nil
}

func (u *UserClient) AddUserRoles(ctx context.Context, tenantId string, userId string, body *user_types.RoleRequest) (*user_types.UserTenant, error) {
	fullUrl := u.cfg.Endpoint.JoinPath(tenantEndpoint, userId, "/roles")

	parsedBody, err := json.Marshal(body)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to marshal create update user roles request body", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fullUrl.String(), bytes.NewBuffer(parsedBody))

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to create update user roles", err)
	}

	req.Header.Set("Content-Type", "application/json")

	var user user_types.UserTenant

	resp, err := u.httpClient.Do(ctx, req, &tenantId)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to update user roles", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to update user roles", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&user)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode update user roles response body", err)
	}

	return &user, nil
}

func (u *UserClient) RemoveUserRoles(ctx context.Context, tenantId string, userId string, body *user_types.RoleRequest) (*user_types.UserTenant, error) {
	fullUrl := u.cfg.Endpoint.JoinPath(tenantEndpoint, userId, "/roles")

	parsedBody, err := json.Marshal(body)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to marshal create update user roles request body", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", fullUrl.String(), bytes.NewBuffer(parsedBody))

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to create update user roles", err)
	}

	req.Header.Set("Content-Type", "application/json")

	var user user_types.UserTenant

	resp, err := u.httpClient.Do(ctx, req, &tenantId)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to update user roles", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to update user roles", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&user)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode update user roles response body", err)
	}

	return &user, nil
}
