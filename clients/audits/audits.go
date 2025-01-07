package audits

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/hightidecrm/frontegg-go/authenticator"
	"github.com/hightidecrm/frontegg-go/clients/audits/audit_types"
	"github.com/hightidecrm/frontegg-go/config"
	"github.com/hightidecrm/frontegg-go/errors"
	"github.com/hightidecrm/frontegg-go/internal/http_client"
)

type AuditsClient struct {
	cfg        *config.FronteggConfig
	auth       *authenticator.FronteggAuthenticator
	httpClient *http_client.FronteggHttpClient
}

// NewAuditsClient creates a new AuditsClient with the given options
func NewAuditsClient(cfg *config.FronteggConfig, auth *authenticator.FronteggAuthenticator, httpClient *http_client.FronteggHttpClient) *AuditsClient {
	return &AuditsClient{cfg: cfg, auth: auth, httpClient: httpClient}
}

// SendAudit sends an audit request to the Frontegg Audits Service
func (a *AuditsClient) SendAudit(ctx context.Context, body *audit_types.AuditRequestParams) error {
	mBody, err := json.Marshal(body)

	if err != nil {
		return err
	}

	reqBody := bytes.NewBuffer(mBody)

	fullUrl := a.cfg.GetAuditsServiceUrl()

	req, err := http.NewRequestWithContext(ctx, "POST", fullUrl.String(), reqBody)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = a.httpClient.Do(ctx, req, &body.TenantID)

	if err != nil {
		return err
	}

	return nil
}

func (a *AuditsClient) GetAudits(ctx context.Context, params *audit_types.AuditRequestParams) (*audit_types.ListAuditsResponse, error) {
	fullUrl := a.cfg.GetAuditsServiceUrl()

	if params != nil {
		values, err := query.Values(params)

		if err != nil {
			return nil, err
		}

		fullUrl.RawQuery = values.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", fullUrl.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	var response audit_types.ListAuditsResponse

	resp, err := a.httpClient.Do(ctx, req, nil)

	if err != nil {
		return nil, err
	}

	if http_client.IsFronteggResponseError(resp) {
		return nil, errors.ErrInvalidResponse
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
