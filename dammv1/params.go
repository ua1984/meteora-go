package dammv1

// SearchParams are query parameters for the SearchPools method.
// Page and Size are required; all other fields are optional.
type SearchParams struct {
	// Page is the 0-based page number to retrieve.
	Page int

	// Size is the number of pools per page.
	Size int

	// Filter is a free-text search string.
	Filter *string

	// SortKey is the field to sort by ("tvl", "volume", "fee_tvl_ratio", "l_m").
	SortKey *string

	// OrderBy is the sort direction ("asc" or "desc").
	OrderBy *string

	// PoolsToTop is a list of pool addresses to prioritize at the top of results.
	PoolsToTop []string

	// Unknown includes pools with unrecognized tokens when true.
	Unknown *bool

	// PoolType filters by pool variant ("dynamic", "multitoken", "lst", "farms").
	PoolType *string

	// IsMonitoring filters to pools under monitoring.
	IsMonitoring *bool

	// HideLowTVL excludes pools whose TVL is below this USD threshold.
	HideLowTVL *float64

	// HideLowAPR excludes pools with low APR when true.
	HideLowAPR *bool

	// IncludeTokenMints is an allowlist of token mint addresses to include.
	IncludeTokenMints []string

	// IncludePoolTokenPairs is an allowlist of pool token pair combinations to include.
	IncludePoolTokenPairs []string

	// Launchpad filters results to pools associated with the specified launchpad addresses.
	Launchpad []string
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

// ListPoolsParams are optional filter parameters for the ListPools method.
type ListPoolsParams struct {
	// Address filters results to the specified pool addresses.
	Address []string

	// Unknown includes pools with unrecognized tokens when true.
	Unknown *bool

	// PoolType filters by pool variant ("dynamic", "multitoken", "lst", "farms").
	PoolType *string

	// IsMonitoring filters to pools under monitoring.
	IsMonitoring *bool

	// HideLowTVL excludes pools whose TVL is below this USD threshold.
	HideLowTVL *float64

	// HideLowAPR excludes pools with low APR when true.
	HideLowAPR *bool

	// Launchpad filters results to pools associated with the specified launchpad addresses.
	Launchpad []string
}
