package user_types

type CreateUserRequest struct {
	Email               string   `json:"email"`
	Name                *string  `json:"name,omitempty"`
	Password            *string  `json:"password,omitempty"`
	Metadata            *string  `json:"metadata,omitempty"`
	VendorMetadata      *string  `json:"vendorMetadata"`
	RoleIds             []string `json:"roleIds"`
	TenantID            string   `json:"tenantId"`
	ExpirationInSeconds *int     `json:"expirationInSeconds,omitempty"`
	MfaBypass           *bool    `json:"mfaBypass,omitempty"`
	SkipInviteEmail     *bool    `json:"skipInviteEmail,omitempty"`
}
