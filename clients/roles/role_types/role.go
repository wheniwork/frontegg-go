package role_types

import "time"

type Role struct {
	ID            string    `json:"id"`
	VendorID      string    `json:"vendorId"`
	TenantID      *string   `json:"tenantId,omitempty"`
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
