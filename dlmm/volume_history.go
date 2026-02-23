package dlmm

// VolumeHistory represents a single volume data point in a pool's history.
type VolumeHistory struct {
	// Timestamp is the Unix timestamp for the start of the period.
	Timestamp int64 `json:"timestamp"`

	// TimestampStr is the human-readable timestamp string.
	TimestampStr string `json:"timestamp_str"`

	// Volume is the total trading volume during the period in USD.
	Volume float64 `json:"volume"`

	// Fees is the total trading fees collected during the period in USD.
	Fees float64 `json:"fees"`

	// ProtocolFees is the protocol's share of fees during the period in USD.
	ProtocolFees float64 `json:"protocol_fees"`
}

// VolumeHistoryResponse wraps volume history data with time range metadata.
type VolumeHistoryResponse struct {
	// StartTime is the Unix timestamp of the earliest data point in the response.
	StartTime int64 `json:"start_time"`

	// EndTime is the Unix timestamp of the latest data point in the response.
	EndTime int64 `json:"end_time"`

	// Timeframe is the resolution of each data point (e.g., "1h", "1d").
	Timeframe string `json:"timeframe"`

	// Data contains the volume history data points.
	Data []VolumeHistory `json:"data"`
}

// VolumeHistoryParams are optional query parameters for the GetVolumeHistory method.
type VolumeHistoryParams struct {
	TimeframeBasedParams
}
