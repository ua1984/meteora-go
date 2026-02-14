package dlmm

// LegacyPair represents a pool from the legacy DLMM API (/pair/all endpoint).
// This endpoint returns a flat structure with string-encoded numeric fields,
// unlike the datapi which uses nested objects and native numbers.
type LegacyPair struct {
	// Address is the on-chain address of the pool.
	Address string `json:"address"`

	// Name is the human-readable pool name (e.g., "SOL-USDC").
	Name string `json:"name"`

	// MintX is the mint address of the first token in the pair.
	MintX string `json:"mint_x"`

	// MintY is the mint address of the second token in the pair.
	MintY string `json:"mint_y"`

	// ReserveX is the reserve account address for token X.
	ReserveX string `json:"reserve_x"`

	// ReserveY is the reserve account address for token Y.
	ReserveY string `json:"reserve_y"`

	// ReserveXAmount is the amount of token X held in reserve (in native units).
	ReserveXAmount int64 `json:"reserve_x_amount"`

	// ReserveYAmount is the amount of token Y held in reserve (in native units).
	ReserveYAmount int64 `json:"reserve_y_amount"`

	// BinStep is the price bin width in basis points.
	BinStep int `json:"bin_step"`

	// BaseFeePercentage is the base swap fee as a string-encoded percentage (e.g., "2").
	BaseFeePercentage string `json:"base_fee_percentage"`

	// MaxFeePercentage is the maximum swap fee as a string-encoded percentage.
	MaxFeePercentage string `json:"max_fee_percentage"`

	// ProtocolFeePercentage is the protocol fee share as a string-encoded percentage.
	ProtocolFeePercentage string `json:"protocol_fee_percentage"`

	// Liquidity is the total liquidity as a string-encoded number.
	Liquidity string `json:"liquidity"`

	// RewardMintX is the mint address of the first farming reward token.
	RewardMintX string `json:"reward_mint_x"`

	// RewardMintY is the mint address of the second farming reward token.
	RewardMintY string `json:"reward_mint_y"`

	// Fees24h is the total fees collected in the last 24 hours in USD.
	Fees24h float64 `json:"fees_24h"`

	// TodayFees is the total fees collected today in USD.
	TodayFees float64 `json:"today_fees"`

	// TradeVolume24h is the trading volume in the last 24 hours in USD.
	TradeVolume24h float64 `json:"trade_volume_24h"`

	// CumulativeTradeVolume is the lifetime cumulative trade volume as a string-encoded number.
	CumulativeTradeVolume string `json:"cumulative_trade_volume"`

	// CumulativeFeeVolume is the lifetime cumulative fee volume as a string-encoded number.
	CumulativeFeeVolume string `json:"cumulative_fee_volume"`

	// CurrentPrice is the current price of token X in terms of token Y.
	CurrentPrice float64 `json:"current_price"`

	// APR is the estimated annual percentage rate from trading fees.
	APR float64 `json:"apr"`

	// APY is the estimated annual percentage yield from trading fees (compounded).
	APY float64 `json:"apy"`

	// FarmAPR is the annual percentage rate from farming rewards.
	FarmAPR float64 `json:"farm_apr"`

	// FarmAPY is the annual percentage yield from farming rewards (compounded).
	FarmAPY float64 `json:"farm_apy"`

	// Hide indicates whether the pool should be hidden from the UI.
	Hide bool `json:"hide"`

	// IsBlacklisted indicates whether the pool has been flagged/blacklisted.
	IsBlacklisted bool `json:"is_blacklisted"`

	// Fees contains fee amounts aggregated across time windows.
	Fees LegacyTimeBuckets `json:"fees"`

	// FeeTVLRatio contains fee-to-TVL ratios across time windows.
	FeeTVLRatio LegacyTimeBuckets `json:"fee_tvl_ratio"`

	// Volume contains trading volume aggregated across time windows.
	Volume LegacyTimeBuckets `json:"volume"`

	// Tags is a list of labels associated with the pool.
	Tags []string `json:"tags"`

	// Launchpad identifies the launchpad that created this pool, if any. Nil if none.
	Launchpad *string `json:"launchpad"`

	// IsVerified indicates whether the pool's tokens are verified.
	IsVerified bool `json:"is_verified"`
}
