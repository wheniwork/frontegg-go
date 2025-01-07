package groups

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/hightidecrm/frontegg-go/clients/groups/group_types"
	"github.com/hightidecrm/frontegg-go/errors"
	"github.com/hightidecrm/frontegg-go/internal/http_client"
)

func (g *GroupClient) CreateGroup(ctx context.Context, tenantId string, body *group_types.CreateGroupRequest) (*group_types.Group, error) {
	fullUrl := g.cfg.Endpoint.JoinPath(baseEndpoint)

	parsedBody, err := json.Marshal(body)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to marshal create group request body", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fullUrl.String(), bytes.NewBuffer(parsedBody))

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to create group", err)

	}

	req.Header.Set("Content-Type", "application/json")

	var group group_types.Group

	resp, err := g.httpClient.Do(ctx, req, &tenantId)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to create group", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to create group", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&group)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode create group response body", err)
	}

	return &group, nil
}

func (g *GroupClient) UpdateGroup(ctx context.Context, groupId string, tenantId string, body *group_types.CreateGroupRequest) (*group_types.Group, error) {
	fullUrl := g.cfg.Endpoint.JoinPath(baseEndpoint, groupId)

	parsedBody, err := json.Marshal(body)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to marshal create group request body", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", fullUrl.String(), bytes.NewBuffer(parsedBody))

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to create group", err)

	}

	req.Header.Set("Content-Type", "application/json")

	var group group_types.Group

	resp, err := g.httpClient.Do(ctx, req, &tenantId)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to create group", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to create group", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&group)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode create group response body", err)
	}

	return &group, nil
}

func (g *GroupClient) DeleteGroup(ctx context.Context, groupId string, tenantId string) error {
	fullUrl := g.cfg.Endpoint.JoinPath(baseEndpoint, groupId)

	req, err := http.NewRequestWithContext(ctx, "DELETE", fullUrl.String(), nil)

	if err != nil {
		return errors.NewFronteggError(nil, "Failed to delete role", err)
	}

	resp, err := g.httpClient.Do(ctx, req, &tenantId)

	if err != nil {
		return errors.NewFronteggError(nil, "Failed to delete role", err)

	}

	if http_client.IsFronteggResponseError(resp) {
		return errors.NewFronteggError(resp, "Failed to delete role", nil)
	}

	return nil
}

func (g *GroupClient) AssignRolesToGroup(ctx context.Context, groupId string, tenantId string, body *group_types.RolesRequest) error {
	fullUrl := g.cfg.Endpoint.JoinPath(baseEndpoint, groupId, "/roles")

	parsedBody, err := json.Marshal(body)

	if err != nil {
		return errors.NewFronteggError(nil, "Failed to marshal create update user roles request body", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fullUrl.String(), bytes.NewBuffer(parsedBody))

	if err != nil {
		return errors.NewFronteggError(nil, "Failed to create update user roles", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := g.httpClient.Do(ctx, req, &tenantId)

	if err != nil {
		return errors.NewFronteggError(nil, "Failed to update user roles", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return errors.NewFronteggError(resp, "Failed to update user roles", nil)
	}

	return nil
}

func (g *GroupClient) RemoveRolesFromGroup(ctx context.Context, groupId string, tenantId string, body *group_types.RolesRequest) error {
	fullUrl := g.cfg.Endpoint.JoinPath(baseEndpoint, groupId, "/roles")

	parsedBody, err := json.Marshal(body)

	if err != nil {
		return errors.NewFronteggError(nil, "Failed to marshal create update user roles request body", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", fullUrl.String(), bytes.NewBuffer(parsedBody))

	if err != nil {
		return errors.NewFronteggError(nil, "Failed to create update user roles", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := g.httpClient.Do(ctx, req, &tenantId)

	if err != nil {
		return errors.NewFronteggError(nil, "Failed to update user roles", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return errors.NewFronteggError(resp, "Failed to update user roles", nil)
	}

	return nil
}

func (g *GroupClient) AssignUsersToGroup(ctx context.Context, groupId string, tenantId string, body *group_types.UsersRequest) error {
	fullUrl := g.cfg.Endpoint.JoinPath(baseEndpoint, groupId, "/users")

	parsedBody, err := json.Marshal(body)

	if err != nil {
		return errors.NewFronteggError(nil, "Failed to marshal create update user roles request body", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fullUrl.String(), bytes.NewBuffer(parsedBody))

	if err != nil {
		return errors.NewFronteggError(nil, "Failed to create update user roles", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := g.httpClient.Do(ctx, req, &tenantId)

	if err != nil {
		return errors.NewFronteggError(nil, "Failed to update user roles", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return errors.NewFronteggError(resp, "Failed to update user roles", nil)
	}

	return nil
}

func (g *GroupClient) RemoveUsersFromGroup(ctx context.Context, groupId string, tenantId string, body *group_types.UsersRequest) error {
	fullUrl := g.cfg.Endpoint.JoinPath(baseEndpoint, groupId, "/users")

	parsedBody, err := json.Marshal(body)

	if err != nil {
		return errors.NewFronteggError(nil, "Failed to marshal create update user roles request body", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", fullUrl.String(), bytes.NewBuffer(parsedBody))

	if err != nil {
		return errors.NewFronteggError(nil, "Failed to create update user roles", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := g.httpClient.Do(ctx, req, &tenantId)

	if err != nil {
		return errors.NewFronteggError(nil, "Failed to update user roles", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return errors.NewFronteggError(resp, "Failed to update user roles", nil)
	}

	return nil
}
