package dlmm

// TimeBuckets holds aggregated metrics across standard time windows.
// Used for volume, fees, protocol fees, and fee/TVL ratio in datapi responses.
type TimeBuckets struct {
	// Min30 is the value for the last 30-minute window.
	Min30 float64 `json:"30m"`

	// Hour1 is the value for the last 1-hour window.
	Hour1 float64 `json:"1h"`

	// Hour2 is the value for the last 2-hour window.
	Hour2 float64 `json:"2h"`

	// Hour4 is the value for the last 4-hour window.
	Hour4 float64 `json:"4h"`

	// Hour12 is the value for the last 12-hour window.
	Hour12 float64 `json:"12h"`

	// Hour24 is the value for the last 24-hour window.
	Hour24 float64 `json:"24h"`
}

// LegacyTimeBuckets holds time-bucketed metrics from the legacy /pair/all endpoint.
// The JSON keys differ from TimeBuckets (e.g., "min_30" vs "30m").
type LegacyTimeBuckets struct {
	// Min30 is the value for the last 30-minute window.
	Min30 float64 `json:"min_30"`

	// Hour1 is the value for the last 1-hour window.
	Hour1 float64 `json:"hour_1"`

	// Hour2 is the value for the last 2-hour window.
	Hour2 float64 `json:"hour_2"`

	// Hour4 is the value for the last 4-hour window.
	Hour4 float64 `json:"hour_4"`

	// Hour12 is the value for the last 12-hour window.
	Hour12 float64 `json:"hour_12"`

	// Hour24 is the value for the last 24-hour window.
	Hour24 float64 `json:"hour_24"`
}
