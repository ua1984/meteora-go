package dynamicvault

// APYBreakdown represents an APY value for a specific lending strategy within a vault.
type APYBreakdown struct {
	// StrategyName is the human-readable name of the lending strategy.
	StrategyName string `json:"strategy_name"`

	// Strategy is the on-chain address of the strategy account.
	Strategy string `json:"strategy"`

	// APY is the annual percentage yield for this strategy.
	APY float64 `json:"apy"`
}

// APYState holds the APY breakdown for a Dynamic Vault across three time horizons,
// with per-strategy detail for each.
type APYState struct {
	// ClosestAPY contains the short-term APY breakdown by strategy.
	ClosestAPY []APYBreakdown `json:"closest_apy"`

	// AverageAPY contains the medium-term average APY breakdown by strategy.
	AverageAPY []APYBreakdown `json:"average_apy"`

	// LongAPY contains the long-term APY breakdown by strategy.
	LongAPY []APYBreakdown `json:"long_apy"`
}

// APYEntry represents a single APY data point at a specific timestamp,
// returned by the time-range APY filter endpoint.
type APYEntry struct {
	// APY is the annual percentage yield at this point in time.
	APY float64 `json:"apy"`

	// Timestamp is the Unix timestamp for this data point.
	Timestamp int64 `json:"timestamp"`
}
