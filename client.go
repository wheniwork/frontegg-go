package frontegg

import (
	"github.com/wheniwork/frontegg-go/authenticator"
	"github.com/wheniwork/frontegg-go/clients/audits"
	"github.com/wheniwork/frontegg-go/clients/entitlements"
	"github.com/wheniwork/frontegg-go/clients/events"
	"github.com/wheniwork/frontegg-go/clients/hosted_login"
	"github.com/wheniwork/frontegg-go/clients/identity"
	"github.com/wheniwork/frontegg-go/clients/tenants"
	"github.com/wheniwork/frontegg-go/clients/users"
	"github.com/wheniwork/frontegg-go/config"
	"github.com/wheniwork/frontegg-go/internal/http_client"
)

type Frontegg struct {
	cfg        *config.FronteggConfig
	auth       *authenticator.FronteggAuthenticator
	httpClient *http_client.FronteggHttpClient

	Audits       *audits.AuditsClient
	Entitlements *entitlements.EntitlementsClient
	Events       *events.EventsClient
	HostedLogin  *hosted_login.HostedLoginClient
	Identity     *identity.IdentityClient

	Tenant *tenants.TenantClient
	User   *users.UserClient
}

// NewFrontegg creates a new Frontegg client with the given options
func NewFrontegg(opts ...config.FronteggOption) (*Frontegg, error) {
	cfg, err := config.NewFronteggConfig(opts...)
	if err != nil {
		return nil, err
	}

	auth := authenticator.NewFronteggAuthenticator(cfg)

	httpClient := http_client.NewFronteggHttpClient(auth, cfg)

	return &Frontegg{
		cfg:          cfg,
		auth:         auth,
		httpClient:   httpClient,
		Audits:       audits.NewAuditsClient(cfg, auth, httpClient),
		Entitlements: entitlements.NewEntitlementsClient(cfg, auth, httpClient),
		Events:       events.NewEventsClient(cfg, auth, httpClient),
		HostedLogin:  hosted_login.NewHostedLoginClient(cfg, auth, httpClient),
		Identity:     identity.NewIdentityClient(cfg, auth, httpClient),

		Tenant: tenants.NewTenantClient(cfg, auth, httpClient),
		User:   users.NewUserClient(cfg, auth, httpClient),
	}, nil
}
