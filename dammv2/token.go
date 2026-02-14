package dammv2

// Token represents SPL token metadata associated with a DAMM v2 pool.
type Token struct {
	// Address is the token's mint address on Solana.
	Address string `json:"address"`

	// Name is the human-readable name of the token (e.g., "Solana").
	Name string `json:"name"`

	// Symbol is the token's ticker symbol (e.g., "SOL").
	Symbol string `json:"symbol"`

	// Decimals is the number of decimal places for the token.
	Decimals int `json:"decimals"`

	// IsVerified indicates whether the token has been verified on the registry.
	IsVerified bool `json:"is_verified"`

	// Holders is the number of unique wallet addresses holding this token.
	Holders int64 `json:"holders"`

	// FreezeAuthorityDisabled indicates whether the token's freeze authority has been revoked.
	FreezeAuthorityDisabled bool `json:"freeze_authority_disabled"`

	// TotalSupply is the total supply of the token.
	TotalSupply float64 `json:"total_supply"`

	// Price is the current USD price of the token.
	Price float64 `json:"price"`

	// MarketCap is the token's total market capitalization in USD.
	MarketCap float64 `json:"market_cap"`
}
