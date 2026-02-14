package dammv1

// Pool represents a DAMM v1 (Dynamic AMM v1) pool. Many numeric fields are
// returned as strings from the API to preserve precision for large values.
type Pool struct {
	// PoolAddress is the on-chain address of the pool account.
	PoolAddress string `json:"pool_address"`

	// PoolTokenMints contains the mint addresses of the tokens in the pool.
	PoolTokenMints []string `json:"pool_token_mints"`

	// PoolTokenAmounts contains the token amounts in the pool as string-encoded numbers.
	PoolTokenAmounts []string `json:"pool_token_amounts"`

	// PoolTokenUSDAmounts contains the USD values of each token in the pool as strings.
	PoolTokenUSDAmounts []string `json:"pool_token_usd_amounts"`

	// Vaults contains the vault addresses associated with this pool.
	Vaults []string `json:"vaults"`

	// VaultLPs contains the vault LP token addresses.
	VaultLPs []string `json:"vault_lps"`

	// LPMint is the mint address of the pool's LP token.
	LPMint string `json:"lp_mint"`

	// PoolTVL is the total value locked in the pool in USD as a string-encoded number.
	PoolTVL string `json:"pool_tvl"`

	// FarmTVL is the total value locked in the farm in USD as a string-encoded number.
	FarmTVL string `json:"farm_tvl"`

	// FarmingPool is the farming pool address, nil if no farming is active.
	FarmingPool *string `json:"farming_pool"`

	// FarmingAPY is the farming annual percentage yield as a string-encoded number.
	FarmingAPY string `json:"farming_apy"`

	// IsMonitoring indicates whether the pool is being monitored by the system.
	IsMonitoring bool `json:"is_monitoring"`

	// PoolOrder is the display ordering priority for the pool.
	PoolOrder int `json:"pool_order"`

	// FarmOrder is the display ordering priority for the farm.
	FarmOrder int `json:"farm_order"`

	// PoolVersion is the version of the pool program.
	PoolVersion int `json:"pool_version"`

	// PoolName is the human-readable pool name (e.g., "SOL-USDC").
	PoolName string `json:"pool_name"`

	// LPDecimal is the number of decimal places for the LP token.
	LPDecimal int `json:"lp_decimal"`

	// FarmRewardDurationEnd is the Unix timestamp when farming rewards end.
	FarmRewardDurationEnd int64 `json:"farm_reward_duration_end"`

	// FarmExpire indicates whether the farming program has expired.
	FarmExpire bool `json:"farm_expire"`

	// PoolLPPriceInUSD is the USD price of one LP token as a string-encoded number.
	PoolLPPriceInUSD string `json:"pool_lp_price_in_usd"`

	// TradingVolume is the 24-hour trading volume in USD.
	TradingVolume float64 `json:"trading_volume"`

	// FeeVolume is the 24-hour fee volume in USD.
	FeeVolume float64 `json:"fee_volume"`

	// WeeklyTradingVolume is the 7-day trading volume in USD.
	WeeklyTradingVolume float64 `json:"weekly_trading_volume"`

	// WeeklyFeeVolume is the 7-day fee volume in USD.
	WeeklyFeeVolume float64 `json:"weekly_fee_volume"`

	// YieldVolume is the yield volume as a string-encoded number.
	YieldVolume string `json:"yield_volume"`

	// AccumulatedTradingVolume is the lifetime trading volume as a string-encoded number.
	AccumulatedTradingVolume string `json:"accumulated_trading_volume"`

	// AccumulatedFeeVolume is the lifetime fee volume as a string-encoded number.
	AccumulatedFeeVolume string `json:"accumulated_fee_volume"`

	// AccumulatedYieldVolume is the lifetime yield volume as a string-encoded number.
	AccumulatedYieldVolume string `json:"accumulated_yield_volume"`

	// TradeAPY is the trading APY as a string-encoded number.
	TradeAPY string `json:"trade_apy"`

	// WeeklyTradeAPY is the 7-day trading APY as a string-encoded number.
	WeeklyTradeAPY string `json:"weekly_trade_apy"`

	// DailyBaseAPY is the daily base APY as a string-encoded number.
	DailyBaseAPY string `json:"daily_base_apy"`

	// WeeklyBaseAPY is the weekly base APY as a string-encoded number.
	WeeklyBaseAPY string `json:"weekly_base_apy"`

	// APR is the annual percentage rate.
	APR float64 `json:"apr"`

	// FarmNew indicates whether the farm was recently created.
	FarmNew bool `json:"farm_new"`

	// Permissioned indicates whether the pool requires permission to join.
	Permissioned bool `json:"permissioned"`

	// Unknown indicates whether the pool has unrecognized tokens.
	Unknown bool `json:"unknown"`

	// TotalFeePct is the total fee percentage as a string-encoded number.
	TotalFeePct string `json:"total_fee_pct"`

	// IsLST indicates whether the pool contains a liquid staking token.
	IsLST bool `json:"is_lst"`

	// IsForex indicates whether the pool is a forex (stablecoin) pair.
	IsForex bool `json:"is_forex"`

	// CreatedAt is the Unix timestamp when the pool was created.
	CreatedAt int64 `json:"created_at"`

	// IsMeme indicates whether the pool contains a meme token.
	IsMeme bool `json:"is_meme"`

	// PoolType identifies the pool variant (e.g., "constant_product", "stable").
	PoolType string `json:"pool_type"`
}
