package stake2earn

// FilterParams are optional query parameters for the FilterVaults method.
type FilterParams struct {
	// Page is the 1-based page number to retrieve.
	Page *int `json:"page,omitempty"`

	// Size is the number of vaults per page.
	Size *int `json:"size,omitempty"`

	// SortBy is the field name to sort results by (e.g., "total_staked_amount_usd").
	SortBy *string `json:"sort_by,omitempty"`

	// SortOrder is the sort direction ("asc" or "desc").
	SortOrder *string `json:"sort_order,omitempty"`

	// SearchTerm filters vaults by token symbol or address substring.
	SearchTerm *string `json:"search_term,omitempty"`
}
