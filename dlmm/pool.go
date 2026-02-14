package dlmm

// Pool represents a DLMM (Dynamic Liquidity Market Maker) pool
// from the datapi endpoint. DLMM pools use a concentrated liquidity
// bin-based model for efficient trading.
type Pool struct {
	// Address is the on-chain address of the pool account.
	Address string `json:"address"`

	// Name is the human-readable pool name (e.g., "SOL-USDC").
	Name string `json:"name"`

	// TokenX is the metadata for the first token in the pair.
	TokenX Token `json:"token_x"`

	// TokenY is the metadata for the second token in the pair.
	TokenY Token `json:"token_y"`

	// ReserveX is the reserve account address for token X.
	ReserveX string `json:"reserve_x"`

	// ReserveY is the reserve account address for token Y.
	ReserveY string `json:"reserve_y"`

	// TokenXAmount is the amount of token X in the pool.
	TokenXAmount float64 `json:"token_x_amount"`

	// TokenYAmount is the amount of token Y in the pool.
	TokenYAmount float64 `json:"token_y_amount"`

	// CreatedAt is the Unix timestamp when the pool was created.
	CreatedAt int64 `json:"created_at"`

	// RewardMintX is the mint address of the first farming reward token, if any.
	RewardMintX string `json:"reward_mint_x"`

	// RewardMintY is the mint address of the second farming reward token, if any.
	RewardMintY string `json:"reward_mint_y"`

	// PoolConfig holds the pool's configuration parameters (bin step, fees).
	PoolConfig PoolConfig `json:"pool_config"`

	// DynamicFeePct is the current dynamic fee percentage, adjusted based on volatility.
	DynamicFeePct float64 `json:"dynamic_fee_pct"`

	// TVL is the total value locked in the pool in USD.
	TVL float64 `json:"tvl"`

	// CurrentPrice is the current price of token X in terms of token Y.
	CurrentPrice float64 `json:"current_price"`

	// APR is the estimated annual percentage rate from trading fees.
	APR float64 `json:"apr"`

	// APY is the estimated annual percentage yield from trading fees (compounded).
	APY float64 `json:"apy"`

	// HasFarm indicates whether the pool has an active farming program.
	HasFarm bool `json:"has_farm"`

	// FarmAPR is the annual percentage rate from farming rewards.
	FarmAPR float64 `json:"farm_apr"`

	// FarmAPY is the annual percentage yield from farming rewards (compounded).
	FarmAPY float64 `json:"farm_apy"`

	// Volume contains trading volume aggregated across time windows.
	Volume TimeBuckets `json:"volume"`

	// Fees contains trading fees aggregated across time windows.
	Fees TimeBuckets `json:"fees"`

	// ProtocolFees contains protocol fees aggregated across time windows.
	ProtocolFees TimeBuckets `json:"protocol_fees"`

	// FeeTVLRatio contains fee-to-TVL ratios across time windows.
	FeeTVLRatio TimeBuckets `json:"fee_tvl_ratio"`

	// CumulativeMetrics holds lifetime cumulative volume and fee data.
	CumulativeMetrics CumulativeMetrics `json:"cumulative_metrics"`

	// IsBlacklisted indicates whether the pool has been flagged/blacklisted.
	IsBlacklisted bool `json:"is_blacklisted"`

	// Launchpad identifies the launchpad that created this pool, if any.
	Launchpad string `json:"launchpad"`

	// Tags is a list of labels associated with the pool (e.g., "memecoin").
	Tags []string `json:"tags"`
}
