package dammv1

// MetricsBreakdown holds aggregate metrics for a specific pool category.
type MetricsBreakdown struct {
	// TVL is the total value locked in USD for this category.
	TVL float64 `json:"tvl"`

	// DailyVolume is the 24-hour trading volume in USD.
	DailyVolume float64 `json:"daily_volume"`

	// TotalVolume is the lifetime cumulative trading volume in USD.
	TotalVolume float64 `json:"total_volume"`

	// DailyFee is the 24-hour fee volume in USD.
	DailyFee float64 `json:"daily_fee"`

	// TotalFee is the lifetime cumulative fee volume in USD.
	TotalFee float64 `json:"total_fee"`
}

// PoolMetrics holds protocol-level aggregate metrics for all DAMM v1 pools,
// with both flat top-level fields and structured breakdowns by category.
type PoolMetrics struct {
	// DynamicAMMTVL is the total TVL in USD for dynamic AMM pools.
	DynamicAMMTVL float64 `json:"dynamic_amm_tvl"`

	// DynamicAMMDailyVolume is the 24h trading volume in USD for dynamic AMM pools.
	DynamicAMMDailyVolume float64 `json:"dynamic_amm_daily_volume"`

	// DynamicAMMTotalVolume is the lifetime volume in USD for dynamic AMM pools.
	DynamicAMMTotalVolume float64 `json:"dynamic_amm_total_volume"`

	// DynamicAMMDailyFee is the 24h fee volume in USD for dynamic AMM pools.
	DynamicAMMDailyFee float64 `json:"dynamic_amm_daily_fee"`

	// DynamicAMMTotalFee is the lifetime fee volume in USD for dynamic AMM pools.
	DynamicAMMTotalFee float64 `json:"dynamic_amm_total_fee"`

	// MultitokensTVL is the total TVL in USD for multitoken pools.
	MultitokensTVL float64 `json:"multitokens_tvl"`

	// MultitokensDailyVolume is the 24h trading volume in USD for multitoken pools.
	MultitokensDailyVolume float64 `json:"multitokens_daily_volume"`

	// MultitokensTotalVolume is the lifetime volume in USD for multitoken pools.
	MultitokensTotalVolume float64 `json:"multitokens_total_volume"`

	// MultitokensDailyFee is the 24h fee volume in USD for multitoken pools.
	MultitokensDailyFee float64 `json:"multitokens_daily_fee"`

	// MultitokensTotalFee is the lifetime fee volume in USD for multitoken pools.
	MultitokensTotalFee float64 `json:"multitokens_total_fee"`

	// DynamicAMM holds structured metrics for dynamic AMM pools.
	DynamicAMM MetricsBreakdown `json:"dynamic_amm"`

	// LST holds structured metrics for liquid staking token pools.
	LST MetricsBreakdown `json:"lst"`

	// Farms holds structured metrics for farming pools.
	Farms MetricsBreakdown `json:"farms"`

	// Multitokens holds structured metrics for multitoken pools.
	Multitokens MetricsBreakdown `json:"multitokens"`
}
