package dammv2

// PoolGroup represents a group of DAMM v2 pools that share the same token pair.
type PoolGroup struct {
	// LexicalOrderMints is the lexicographically sorted concatenation of both
	// token mint addresses, used as a unique identifier for the group.
	LexicalOrderMints string `json:"lexical_order_mints"`

	// GroupName is the human-readable name for the group (e.g., "SOL-USDC").
	GroupName string `json:"group_name"`

	// TokenX is the mint address of the first token in the pair.
	TokenX string `json:"token_x"`

	// TokenY is the mint address of the second token in the pair.
	TokenY string `json:"token_y"`

	// PoolCount is the number of pools in this group.
	PoolCount int `json:"pool_count"`

	// TotalTVL is the combined TVL in USD across all pools in the group.
	TotalTVL float64 `json:"total_tvl"`

	// TotalVolume is the combined 24h trading volume in USD across all pools.
	TotalVolume float64 `json:"total_volume"`

	// MaxFeeTVLRatio is the highest fee/TVL ratio among pools in the group.
	MaxFeeTVLRatio float64 `json:"max_fee_tvl_ratio"`

	// HasFarm indicates whether any pool in the group has an active farm.
	HasFarm bool `json:"has_farm"`

	// MaxFarmAPR is the highest farming APR among pools in the group.
	MaxFarmAPR float64 `json:"max_farm_apr"`
}
