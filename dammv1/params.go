package dammv1

// SearchParams are optional query parameters for the SearchPools method.
type SearchParams struct {
	// Page is the 1-based page number to retrieve.
	Page *int `json:"page,omitempty"`

	// Size is the number of pools per page.
	Size *int `json:"size,omitempty"`

	// SearchTerm filters pools by name, symbol, or address substring.
	SearchTerm *string `json:"search_term,omitempty"`

	// SortBy is the field name to sort results by.
	SortBy *string `json:"sort_by,omitempty"`

	// SortOrder is the sort direction ("asc" or "desc").
	SortOrder *string `json:"sort_order,omitempty"`
}

// PaginationParams are query parameters for paginated endpoints (e.g., ListPoolsWithFarm).
type PaginationParams struct {
	// Page is the 1-based page number to retrieve.
	Page *int `json:"page,omitempty"`

	// Size is the number of records per page.
	Size *int `json:"size,omitempty"`
}

// AlphaVaultParams are optional filter parameters for the ListAlphaVaults method.
type AlphaVaultParams struct {
	// VaultAddress filters results to the specified vault addresses.
	VaultAddress []string

	// PoolAddress filters results to vaults associated with the specified pool addresses.
	PoolAddress []string

	// BaseMint filters results to vaults with the specified base mint addresses.
	BaseMint []string
}
