package frontegg

import (
	"github.com/hightidecrm/frontegg/authenticator"
	"github.com/hightidecrm/frontegg/clients/audits"
	"github.com/hightidecrm/frontegg/clients/entitlements"
	"github.com/hightidecrm/frontegg/clients/events"
	"github.com/hightidecrm/frontegg/clients/hosted_login"
	"github.com/hightidecrm/frontegg/clients/identity"
	"github.com/hightidecrm/frontegg/clients/tenants"
	"github.com/hightidecrm/frontegg/clients/users"
	"github.com/hightidecrm/frontegg/config"
	"github.com/hightidecrm/frontegg/internal/http_client"
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
