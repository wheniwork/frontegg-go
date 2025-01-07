package tenants

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/hightidecrm/frontegg-go/clients/tenants/tenant_types"
	"github.com/hightidecrm/frontegg-go/errors"
	"github.com/hightidecrm/frontegg-go/internal/http_client"
	"github.com/hightidecrm/frontegg-go/types"
)

func (t *TenantClient) GetTenants(ctx context.Context, params *types.QueryParams) (*types.FronteggPagination[tenant_types.Tenant], error) {
	fullUrl := t.cfg.Endpoint.JoinPath(baseEndpoint)

	if params != nil {
		parsedParams, err := query.Values(params)

		if err != nil {
			return nil, errors.NewFronteggError(nil, "Failed to parse query params", err)
		}

		fullUrl.RawQuery = parsedParams.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", fullUrl.String(), nil)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to get tenants", err)
	}

	var tenants types.FronteggPagination[tenant_types.Tenant]

	resp, err := t.httpClient.Do(ctx, req, nil)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to get tenants", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to get tenants", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&tenants)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode get tenants response body", err)
	}

	return nil, nil
}

func (t *TenantClient) GetTenant(ctx context.Context, tenantId string) (*tenant_types.Tenant, error) {
	fullUrl := t.cfg.Endpoint.JoinPath(baseEndpoint, tenantId)

	req, err := http.NewRequestWithContext(ctx, "GET", fullUrl.String(), nil)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to get tenant", err)
	}

	var tenants tenant_types.Tenant

	resp, err := t.httpClient.Do(ctx, req, nil)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to get tenant", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to get tenant", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&tenants)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode get tenant response body", err)
	}

	return nil, nil
}
