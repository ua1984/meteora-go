package dammv1

// FeeConfig represents a fee configuration entry for a DAMM v1 pool.
type FeeConfig struct {
	// ConfigAddress is the on-chain address of this fee configuration.
	ConfigAddress string `json:"config_address"`

	// CreatorAuthority is the public key of the authority that created this config.
	CreatorAuthority string `json:"creator_authority"`

	// ActivateDurationAfterTradeInSeconds is the cooldown duration in seconds
	// after a trade before the fee configuration change takes effect.
	ActivateDurationAfterTradeInSeconds int `json:"activate_duration_after_trade_in_seconds"`

	// FeePercentage is the fee percentage as a string-encoded number.
	FeePercentage string `json:"fee_percentage"`
}
