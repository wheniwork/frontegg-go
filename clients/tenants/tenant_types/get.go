package tenant_types

import "time"

type Tenant struct {
	VendorID   string                 `json:"vendorId"`
	TenantID   string                 `json:"tenantId"`
	Name       string                 `json:"name"`
	UpdatedAt  time.Time              `json:"updatedAt"`
	CreatedAt  time.Time              `json:"createdAt"`
	DeletedAt  *time.Time             `json:"deletedAt,omitempty"`
	IsReseller bool                   `json:"isReseller"`
	Logo       *string                `json:"logo,omitempty"`
	LogoURL    *string                `json:"logoUrl,omitempty"`
	Domain     string                 `json:"domain"`
	Metadata   map[string]interface{} `json:"metadata"`
}
