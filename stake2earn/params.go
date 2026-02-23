package stake2earn

// FilterParams are optional query parameters for the FilterVaults method.
type FilterParams struct {
	// PoolAddresses filters vaults by pool address. Maximum of 100 addresses.
	PoolAddresses []string
}
