package dlmm

// PoolConfig holds the configuration parameters for a DLMM pool.
type PoolConfig struct {
	// BinStep is the price bin width in basis points. A smaller bin step provides
	// finer price granularity but requires more bins to cover the same price range.
	BinStep int `json:"bin_step"`

	// BaseFeePct is the base swap fee percentage charged on each trade.
	BaseFeePct float64 `json:"base_fee_pct"`

	// MaxFeePct is the maximum swap fee percentage that can be charged,
	// including dynamic fee adjustments based on volatility.
	MaxFeePct float64 `json:"max_fee_pct"`

	// ProtocolFeePct is the percentage of swap fees allocated to the protocol.
	ProtocolFeePct float64 `json:"protocol_fee_pct"`
}
