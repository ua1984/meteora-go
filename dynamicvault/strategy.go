package dynamicvault

// Strategy represents a lending strategy used by a Dynamic Vault to generate yield.
// Vaults allocate tokens across multiple strategies to optimize returns.
type Strategy struct {
	// Pubkey is the on-chain address of the strategy account.
	Pubkey string `json:"pubkey"`

	// Reserve is the reserve account address used by the strategy.
	Reserve string `json:"reserve"`

	// StrategyType identifies the lending protocol (e.g., "solend", "mango", "marginfi").
	StrategyType string `json:"strategy_type"`

	// StrategyName is the human-readable name of the strategy.
	StrategyName string `json:"strategy_name"`

	// Liquidity is the amount of tokens currently allocated to this strategy in native units.
	Liquidity int64 `json:"liquidity"`

	// MaxAllocation is the maximum percentage of vault funds that can be allocated to this strategy.
	MaxAllocation float64 `json:"max_allocation"`

	// Isolated indicates whether this strategy is isolated from the main vault pool.
	Isolated bool `json:"isolated"`

	// Disabled indicates whether this strategy is currently disabled.
	Disabled bool `json:"disabled"`

	// SafeUtilizationThreshold is the maximum utilization percentage considered safe for this strategy.
	SafeUtilizationThreshold float64 `json:"safe_utilization_threshold"`
}
