package tenants

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/wheniwork/frontegg-go/clients/tenants/tenant_types"
	"github.com/wheniwork/frontegg-go/errors"
	"github.com/wheniwork/frontegg-go/internal/http_client"
)

func (t *TenantClient) CreateTenant(ctx context.Context, body *tenant_types.CreateTenantRequest) (*tenant_types.Tenant, error) {
	fullUrl := t.cfg.Endpoint.JoinPath(v1BaseEndpoint)

	parsedBody, err := json.Marshal(body)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to marshal create tenant request body", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fullUrl.String(), bytes.NewBuffer(parsedBody))

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to create tenant", err)
	}

	req.Header.Set("Content-Type", "application/json")

	var tenant tenant_types.Tenant

	resp, err := t.httpClient.Do(ctx, req, nil)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to create tenant", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to create tenant", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&tenant)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode create tenant response body", err)
	}

	return &tenant, nil
}

func (t *TenantClient) UpdateTenant(ctx context.Context, tenantId string, body *tenant_types.UpdateTenantRequest) (*tenant_types.Tenant, error) {
	fullUrl := t.cfg.Endpoint.JoinPath(baseEndpoint, tenantId)

	parsedBody, err := json.Marshal(body)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to marshal update tenant request body", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", fullUrl.String(), bytes.NewBuffer(parsedBody))

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to update tenant", err)
	}

	req.Header.Set("Content-Type", "application/json")

	var tenant tenant_types.Tenant

	resp, err := t.httpClient.Do(ctx, req, nil)

	if err != nil {
		return nil, errors.NewFronteggError(nil, "Failed to update tenant", err)
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.NewFronteggError(resp, "Failed to update tenant", nil)
	}

	err = json.NewDecoder(resp.Body).Decode(&tenant)

	if err != nil {
		return nil, errors.NewFronteggError(resp, "Failed to decode update tenant response body", err)
	}

	return &tenant, nil
}
