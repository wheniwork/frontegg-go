package tenant_types

type UpdateTenantRequest struct {
	Name           string  `json:"name,omitempty"`
	Status         *string `json:"status,omitempty"`
	Website        *string `json:"website,omitempty"`
	ApplicationURL *string `json:"applicationUrl,omitempty"`
	Logo           *string `json:"logo,omitempty"`
	LogoURL        *string `json:"logoUrl,omitempty"`
	Address        *string `json:"address,omitempty"`
	Timezone       *string `json:"timezone,omitempty"`
	Currency       *string `json:"currency,omitempty"`
	CreatorName    *string `json:"creatorName,omitempty"`
	CreatorEmail   *string `json:"creatorEmail,omitempty"`
	IsReseller     bool    `json:"isReseller,omitempty"`
}
