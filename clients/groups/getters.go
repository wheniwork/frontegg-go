package groups

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/hightidecrm/frontegg-go/clients/groups/group_types"
	"github.com/hightidecrm/frontegg-go/errors"
	"github.com/hightidecrm/frontegg-go/internal/http_client"
	"github.com/hightidecrm/frontegg-go/types"
)

func (g *GroupClient) GetGroups(ctx context.Context, tenantId *string, params *types.QueryParams) (*types.FronteggPagination[group_types.Group], error) {
	fullUrl := g.cfg.Endpoint.JoinPath(v2BaseEndpoint)

	if params != nil {
		parsedParams, err := query.Values(params)

		if err != nil {
			return nil, errors.NewFronteggError(nil, "Failed to parse query params", err)
		}

		fullUrl.RawQuery = parsedParams.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", fullUrl.String(), nil)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to get groups", err)
	}

	var groups types.FronteggPagination[group_types.Group]

	resp, err := g.httpClient.Do(ctx, req, nil)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to get groups", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to get groups", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&groups)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode get groups response body", err)
	}

	return &groups, nil
}

func (g *GroupClient) GetGroup(ctx context.Context, groupId string, tenantId *string) (*group_types.Group, error) {
	fullUrl := g.cfg.Endpoint.JoinPath(baseEndpoint, groupId)

	req, err := http.NewRequestWithContext(ctx, "GET", fullUrl.String(), nil)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to get group", err)
	}

	var group group_types.Group

	resp, err := g.httpClient.Do(ctx, req, tenantId)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to get group", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to get group", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&group)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode get group response body", err)
	}

	return &group, nil
}
