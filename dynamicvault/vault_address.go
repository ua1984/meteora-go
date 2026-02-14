package dynamicvault

// VaultAddress holds the key addresses associated with a Dynamic Vault.
type VaultAddress struct {
	// Symbol is the ticker symbol of the vault's underlying token (e.g., "SOL", "USDC").
	Symbol string `json:"symbol"`

	// VaultAddress is the on-chain address of the vault account.
	VaultAddress string `json:"vault_address"`

	// LPMintAddress is the mint address of the vault's LP token.
	LPMintAddress string `json:"lp_mint_address"`

	// FeeAddress is the address of the fee collection account.
	FeeAddress string `json:"fee_address"`
}
