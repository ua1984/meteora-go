package dammv2

// OHLCV represents a single candlestick (Open-High-Low-Close-Volume) data point
// for a DAMM v2 pool's price history.
type OHLCV struct {
	// Timestamp is the Unix timestamp for the start of the candle period.
	Timestamp int64 `json:"timestamp"`

	// TimestampStr is the human-readable timestamp string.
	TimestampStr string `json:"timestamp_str"`

	// Open is the opening price at the start of the period.
	Open float64 `json:"open"`

	// High is the highest price during the period.
	High float64 `json:"high"`

	// Low is the lowest price during the period.
	Low float64 `json:"low"`

	// Close is the closing price at the end of the period.
	Close float64 `json:"close"`

	// Volume is the total trading volume during the period in USD.
	Volume float64 `json:"volume"`
}

// OHLCVResponse wraps OHLCV candlestick data with time range metadata.
type OHLCVResponse struct {
	// StartTime is the Unix timestamp of the earliest candle in the response.
	StartTime int64 `json:"start_time"`

	// EndTime is the Unix timestamp of the latest candle in the response.
	EndTime int64 `json:"end_time"`

	// Timeframe is the resolution of each candle (e.g., "1m", "15m", "1h", "1d").
	Timeframe string `json:"timeframe"`

	// Data contains the OHLCV candle data points.
	Data []OHLCV `json:"data"`
}

// OHLCVParams are optional query parameters for the GetOHLCV method.
type OHLCVParams struct {
	// Timeframe is the candle interval. Allowed values: 5m 30m 1h 2h 4h 12h 24h.
	// If omitted, the API uses 24h.
	Timeframe *string `json:"timeframe,omitempty"`

	// StartTime is the Unix timestamp in seconds (inclusive).
	// If omitted, the API uses a default range based on Timeframe.
	StartTime *int64 `json:"start_time,omitempty"`

	// EndTime is the Unix timestamp in seconds (inclusive).
	// If omitted, the API uses "now" as the end.
	EndTime *int64 `json:"end_time,omitempty"`
}
