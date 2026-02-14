package dynamicvault

// VirtualPrice represents a virtual price data point for a vault strategy.
// The virtual price reflects the accumulated yield over time — a price
// above 1.0 indicates positive returns.
type VirtualPrice struct {
	// Price is the virtual price as a string-encoded decimal (e.g., "1.00355").
	Price string `json:"price"`

	// Timestamp is the Unix timestamp for this data point.
	Timestamp int64 `json:"timestamp"`
}
