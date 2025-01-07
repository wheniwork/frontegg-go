package roles

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/hightidecrm/frontegg/clients/roles/role_types"
	"github.com/hightidecrm/frontegg/errors"
	"github.com/hightidecrm/frontegg/internal/http_client"
)

func (r *RoleClient) CreateRole(ctx context.Context, tenantId string, body *role_types.CreateRoleRequest) (*role_types.Role, error) {
	fullUrl := r.cfg.Endpoint.JoinPath(baseEndpoint)

	parsedBody, err := json.Marshal(body)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to marshal create role request body", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fullUrl.String(), bytes.NewBuffer(parsedBody))

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to create role", err)

	}

	req.Header.Set("Content-Type", "application/json")

	var role role_types.Role

	resp, err := r.httpClient.Do(ctx, req, &tenantId)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to create role", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to create role", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&role)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode create role response body", err)
	}

	return &role, nil
}

func (r *RoleClient) UpdateRole(ctx context.Context, roleId string, tenantId string, body *role_types.UpdateRoleRequest) (*role_types.Role, error) {
	fullUrl := r.cfg.Endpoint.JoinPath(v1BaseEndpoint, roleId)

	parsedBody, err := json.Marshal(body)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to marshal create role request body", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fullUrl.String(), bytes.NewBuffer(parsedBody))

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to create role", err)

	}

	req.Header.Set("Content-Type", "application/json")

	var role role_types.Role

	resp, err := r.httpClient.Do(ctx, req, &tenantId)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to create role", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to create role", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&role)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode create role response body", err)
	}

	return &role, nil
}

func (r *RoleClient) SetRolePermissions(ctx context.Context, roleId string, tenantId string, body *role_types.AssignPermissionsToRoleRequest) (*role_types.Role, error) {
	fullUrl := r.cfg.Endpoint.JoinPath(v1BaseEndpoint, roleId, "/permissions")

	parsedBody, err := json.Marshal(body)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to marshal create role request body", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fullUrl.String(), bytes.NewBuffer(parsedBody))

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to create role", err)

	}

	req.Header.Set("Content-Type", "application/json")

	var role role_types.Role

	resp, err := r.httpClient.Do(ctx, req, &tenantId)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to create role", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to create role", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&role)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode create role response body", err)
	}

	return &role, nil
}

func (r *RoleClient) DeleteRole(ctx context.Context, roleId string, tenantId string) error {
	fullUrl := r.cfg.Endpoint.JoinPath(v1BaseEndpoint, roleId)

	req, err := http.NewRequestWithContext(ctx, "DELETE", fullUrl.String(), nil)

	if err != nil {
		return errors.NewFronteggError(nil, "Failed to delete role", err)
	}

	resp, err := r.httpClient.Do(ctx, req, &tenantId)

	if err != nil {
		return errors.NewFronteggError(nil, "Failed to delete role", err)

	}

	if http_client.IsFronteggResponseError(resp) {
		return errors.NewFronteggError(resp, "Failed to delete role", nil)
	}

	return nil
}
