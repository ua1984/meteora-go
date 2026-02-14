package dynamicvault

// VaultInfo holds detailed information about a Dynamic Vault, including
// its current state, APY metrics, and associated strategies.
type VaultInfo struct {
	// Symbol is the ticker symbol of the vault's underlying token (e.g., "SOL", "USDC").
	Symbol string `json:"symbol"`

	// TokenAddress is the mint address of the vault's underlying token.
	TokenAddress string `json:"token_address"`

	// Pubkey is the on-chain address of the vault account.
	Pubkey string `json:"pubkey"`

	// IsMonitoring indicates whether the vault is actively monitored.
	IsMonitoring bool `json:"is_monitoring"`

	// VaultOrder is the display ordering priority for the vault.
	VaultOrder int `json:"vault_order"`

	// USDRate is the current USD exchange rate for the vault's token.
	USDRate float64 `json:"usd_rate"`

	// ClosestAPY is the short-term APY percentage based on recent performance.
	ClosestAPY float64 `json:"closest_apy"`

	// AverageAPY is the medium-term average APY percentage.
	AverageAPY float64 `json:"average_apy"`

	// LongAPY is the long-term APY percentage.
	LongAPY float64 `json:"long_apy"`

	// EarnedAmount is the total yield earned by the vault in native token units.
	EarnedAmount int64 `json:"earned_amount"`

	// VirtualPrice is the virtual price of the vault's LP token as a string-encoded decimal.
	// A virtual price > 1.0 indicates accumulated yield.
	VirtualPrice string `json:"virtual_price"`

	// Enabled indicates whether the vault is accepting deposits (1 = enabled, 0 = disabled).
	Enabled int `json:"enabled"`

	// LPMint is the mint address of the vault's LP token.
	LPMint string `json:"lp_mint"`

	// FeePubkey is the address of the fee collection account.
	FeePubkey string `json:"fee_pubkey"`

	// TotalAmount is the total amount of tokens in the vault in native units.
	TotalAmount int64 `json:"total_amount"`

	// TotalAmountWithProfit is the total amount including unrealized profit in native units.
	TotalAmountWithProfit int64 `json:"total_amount_with_profit"`

	// TokenAmount is the amount of the underlying token held directly in the vault.
	TokenAmount int64 `json:"token_amount"`

	// FeeAmount is the accumulated fee amount in native token units.
	FeeAmount int64 `json:"fee_amount"`

	// LPSupply is the total supply of the vault's LP token in native units.
	LPSupply int64 `json:"lp_supply"`

	// EarnedUSDAmount is the total yield earned by the vault in USD.
	EarnedUSDAmount float64 `json:"earned_usd_amount"`

	// Strategies contains the lending strategies used by this vault.
	Strategies []Strategy `json:"strategies"`

	// Timestamp is the Unix timestamp when this data was last updated.
	Timestamp int64 `json:"timestamp"`
}

// VaultState holds the current state of a vault. It has the same structure as VaultInfo.
type VaultState = VaultInfo
