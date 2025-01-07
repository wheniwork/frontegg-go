package types

type SortDirection string

const (
	ASC  SortDirection = "ASC"
	DESC SortDirection = "DESC"
)

type QueryParams struct {
	Limit     int           `json:"limit" url:"_limit,omitempty"`
	Offset    int           `json:"offset" url:"_offset,omitempty"`
	SortBy    string        `json:"sortBy" url:"_sortBy,omitempty"`
	SortDir   SortDirection `json:"sortDir" url:"_order,omitempty"`
	TenantIds []string      `json:"tenantIds" url:"_tenantIds,omitempty"`
}
