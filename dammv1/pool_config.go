package dammv1

// PoolConfig represents a DAMM v1 pool configuration that defines
// the fee structure and activation rules for pools created with this config.
type PoolConfig struct {
	// ConfigAddress is the on-chain address of this configuration account.
	ConfigAddress string `json:"config_address"`

	// TradeFeeBPS is the trading fee in basis points (1 BPS = 0.01%).
	TradeFeeBPS int `json:"trade_fee_bps"`

	// ProtocolFeeBPS is the protocol fee in basis points.
	ProtocolFeeBPS int `json:"protocol_fee_bps"`

	// ActivationDuration is the duration in seconds before trading activates after pool creation.
	ActivationDuration int `json:"activation_duration"`

	// VaultConfigKey is the key for the associated vault configuration.
	VaultConfigKey string `json:"vault_config_key"`

	// PoolCreatorAuthority is the public key of the authority that can create pools with this config.
	PoolCreatorAuthority string `json:"pool_creator_authority"`

	// ActivationType determines when trading begins (0 = immediate, 1 = slot-based, 2 = time-based).
	ActivationType int `json:"activation_type"`
}
