package roles

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/hightidecrm/frontegg-go/clients/roles/role_types"
	"github.com/hightidecrm/frontegg-go/errors"
	"github.com/hightidecrm/frontegg-go/internal/http_client"
	"github.com/hightidecrm/frontegg-go/types"
)

func (r *RoleClient) GetRoles(ctx context.Context, tenantId *string, params *types.QueryParams) (*types.FronteggPagination[role_types.Role], error) {
	fullUrl := r.cfg.Endpoint.JoinPath(baseEndpoint)

	if params != nil {
		parsedParams, err := query.Values(params)

		if err != nil {
			return nil, errors.NewFronteggError(nil, "Failed to parse query params", err)
		}

		fullUrl.RawQuery = parsedParams.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", fullUrl.String(), nil)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to get users", err)
	}

	var tenants types.FronteggPagination[role_types.Role]

	resp, err := r.httpClient.Do(ctx, req, nil)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to get users", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to get users", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&tenants)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode get users response body", err)
	}

	return &tenants, nil
}
