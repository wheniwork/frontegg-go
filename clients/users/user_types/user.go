package user_types

import "time"

type UserRole struct {
	ID            string    `json:"id"`
	VendorID      string    `json:"vendorId"`
	TenantID      string    `json:"tenantId"`
	Key           string    `json:"key"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	IsDefault     bool      `json:"isDefault"`
	FirstUserRole bool      `json:"firstUserRole"`
	Level         int       `json:"level"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Permissions   []string  `json:"permissions"`
}

type UserPermission struct {
	ID           string    `json:"id"`
	Key          string    `json:"key"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	RoleIds      []string  `json:"roleIds"`
	CategoryID   string    `json:"categoryId"`
	FePermission bool      `json:"fePermission"`
}

type UserTenant struct {
	TenantID                string     `json:"tenantId"`
	Roles                   []UserRole `json:"roles"`
	TemporaryExpirationDate time.Time  `json:"temporaryExpirationDate"`
	IsDisabled              bool       `json:"isDisabled"`
}

type User struct {
	ID                      string           `json:"id"`
	Email                   string           `json:"email"`
	Name                    string           `json:"name"`
	ProfilePictureURL       string           `json:"profilePictureUrl"`
	Sub                     string           `json:"sub"`
	Verified                bool             `json:"verified"`
	MfaEnrolled             bool             `json:"mfaEnrolled"`
	MfaBypass               bool             `json:"mfaBypass"`
	PhoneNumber             string           `json:"phoneNumber"`
	Roles                   []UserRole       `json:"roles"`
	Permissions             []UserPermission `json:"permissions"`
	Provider                string           `json:"provider"`
	TenantID                string           `json:"tenantId"`
	TenantIds               []string         `json:"tenantIds"`
	ActivatedForTenant      bool             `json:"activatedForTenant"`
	IsLocked                bool             `json:"isLocked"`
	Tenants                 []UserTenant     `json:"tenants"`
	Invisible               bool             `json:"invisible"`
	SuperUser               bool             `json:"superUser"`
	Metadata                string           `json:"metadata"`
	VendorMetadata          string           `json:"vendorMetadata"`
	CreatedAt               time.Time        `json:"createdAt"`
	LastLogin               time.Time        `json:"lastLogin"`
	Groups                  []any            `json:"groups"`
	SubAccountAccessAllowed bool             `json:"subAccountAccessAllowed"`
	ManagedBy               string           `json:"managedBy"`
}
