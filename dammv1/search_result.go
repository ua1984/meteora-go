package dammv1

// SearchResult wraps pool search results with pagination metadata.
type SearchResult struct {
	// Data contains the pools matching the search query.
	Data []Pool `json:"data"`

	// Page is the current page number.
	Page int `json:"page"`

	// TotalCount is the total number of pools matching the search query.
	TotalCount int `json:"total_count"`
}
