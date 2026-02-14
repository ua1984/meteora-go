package dammv2

// PoolConfig holds the configuration parameters for a DAMM v2 pool.
// DAMM v2 pools have more configuration options than DLMM, including
// concentrated liquidity settings and activation controls.
type PoolConfig struct {
	// CollectFeeMode determines how fees are collected (0 = standard, 1 = quote only).
	CollectFeeMode int `json:"collect_fee_mode"`

	// BaseFeeMode determines the base fee calculation method.
	BaseFeeMode int `json:"base_fee_mode"`

	// BaseFeePct is the base swap fee percentage.
	BaseFeePct float64 `json:"base_fee_pct"`

	// ProtocolFeePct is the percentage of swap fees allocated to the protocol.
	ProtocolFeePct float64 `json:"protocol_fee_pct"`

	// PartnerFeePct is the percentage of swap fees allocated to the partner.
	PartnerFeePct float64 `json:"partner_fee_pct"`

	// ReferralFeePct is the percentage of swap fees allocated to the referrer.
	ReferralFeePct float64 `json:"referral_fee_pct"`

	// DynamicFeeInitialized indicates whether dynamic fee adjustment is enabled.
	DynamicFeeInitialized bool `json:"dynamic_fee_initialized"`

	// PoolType identifies the pool variant (0 = constant product, 1 = stable, etc.).
	PoolType int `json:"pool_type"`

	// ConcentratedLiquidity indicates whether the pool uses concentrated liquidity.
	ConcentratedLiquidity bool `json:"concentrated_liquidity"`

	// MinPrice is the minimum price bound for concentrated liquidity pools.
	MinPrice float64 `json:"min_price"`

	// MaxPrice is the maximum price bound for concentrated liquidity pools.
	MaxPrice float64 `json:"max_price"`

	// ActivationType determines when trading begins (0 = immediate, 1 = slot-based, 2 = time-based).
	ActivationType int `json:"activation_type"`

	// ActivationPoint is the slot number or Unix timestamp when trading activates,
	// depending on ActivationType.
	ActivationPoint int64 `json:"activation_point"`
}
