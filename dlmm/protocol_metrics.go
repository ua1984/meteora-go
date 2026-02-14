package dlmm

// ProtocolMetrics holds protocol-wide aggregate metrics for all DLMM pools.
// Retrieved from the /stats/protocol_metrics endpoint.
type ProtocolMetrics struct {
	// TotalTVL is the total value locked across all DLMM pools in USD.
	TotalTVL float64 `json:"total_tvl"`

	// Volume24h is the total trading volume across all pools in the last 24 hours in USD.
	Volume24h float64 `json:"volume_24h"`

	// Fee24h is the total fees collected across all pools in the last 24 hours in USD.
	Fee24h float64 `json:"fee_24h"`

	// TotalVolume is the lifetime cumulative trading volume across all pools in USD.
	TotalVolume float64 `json:"total_volume"`

	// TotalFees is the lifetime cumulative fees collected across all pools in USD.
	TotalFees float64 `json:"total_fees"`

	// TotalPools is the total number of DLMM pools.
	TotalPools int `json:"total_pools"`
}
