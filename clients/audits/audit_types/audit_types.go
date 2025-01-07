package audit_types

import "github.com/hightidecrm/frontegg-go/types"

type AuditSeverity string

const (
	Info     AuditSeverity = "Info"
	Medium   AuditSeverity = "Medium"
	High     AuditSeverity = "High"
	Critical AuditSeverity = "Critical"
	Error    AuditSeverity = "Error"
)

type Audit struct {
	Severity    AuditSeverity `json:"severity"`
	IP          *string       `json:"ip,omitempty"`
	Email       *string       `json:"email,omitempty"`
	Action      *string       `json:"action,omitempty"`
	Description *string       `json:"description,omitempty"`
	UserAgent   *string       `json:"userAgent,omitempty"`
	VendorID    string        `json:"vendorId"`
	TenantID    string        `json:"tenantId"`
	FranteggID  string        `json:"frontegg_id"`
	CreatedAt   string        `json:"createdAt"`
	UpdatedAt   string        `json:"updatedAt"`
}

type SendAuditParams struct {
	TenantID string                 `json:"tenantId"`
	Severity AuditSeverity          `json:"severity"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

type AuditSortField string

type AuditRequestParams struct {
	TenantID      string               `json:"tenantId" url:"-"`
	Filter        *string              `json:"filter,omitempty" url:"filter,omitempty"`
	SortBy        *AuditSortField      `json:"sortBy,omitempty" url:"sortBy,omitempty"`
	SortDirection *types.SortDirection `json:"sortDirection,omitempty" url:"sortDirection,omitempty"`
	Offset        int                  `json:"offset" url:"offset"`
	Count         int                  `json:"count" url:"count"`
	Filters       interface{}          `json:"filters,omitempty" url:"filters,omitempty"`
}

type GetAuditStatsParams struct {
	TenantID string      `json:"tenantId" url:"-"`
	Filters  interface{} `json:"filters,omitempty" url:"filters,omitempty"`
}

type ListAuditsResponse struct {
	Data  []Audit `json:"data"`
	Total int     `json:"total"`
}
