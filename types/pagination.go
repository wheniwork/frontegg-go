package types

type FronteggPagination[T any] struct {
	Metadata struct {
		TotalItems int `json:"totalItems"`
		TotalPages int `json:"totalPages"`
	} `json:"_metadata"`
	Links struct {
		Prev  string `json:"prev"`
		Next  string `json:"next"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"_links"`
	Items []T `json:"items"`
}
