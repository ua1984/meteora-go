package dammv2

// PaginatedResponse is a generic wrapper for paginated API responses
// from the DAMM v2 datapi endpoints.
type PaginatedResponse[T any] struct {
	// Total is the total number of records matching the query across all pages.
	Total int `json:"total"`

	// Pages is the total number of pages available.
	Pages int `json:"pages"`

	// CurrentPage is the 1-based index of the current page.
	CurrentPage int `json:"current_page"`

	// PageSize is the number of records per page.
	PageSize int `json:"page_size"`

	// Data contains the records for the current page.
	Data []T `json:"data"`
}
