package dlmm

// CumulativeMetrics holds lifetime cumulative metrics for a DLMM pool.
type CumulativeMetrics struct {
	// Volume is the total cumulative trading volume in USD.
	Volume float64 `json:"volume"`

	// TradeFee is the total cumulative trade fees earned in USD.
	TradeFee float64 `json:"trade_fee"`

	// ProtocolFee is the total cumulative protocol fees collected in USD.
	ProtocolFee float64 `json:"protocol_fee"`
}
