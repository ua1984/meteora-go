package dammv1

// AlphaVault represents an alpha vault used for token launches via DAMM v1 pools.
// Alpha vaults allow projects to bootstrap liquidity through prorata or FCFS mechanisms.
type AlphaVault struct {
	// VaultAddress is the on-chain address of the alpha vault.
	VaultAddress string `json:"vault_address"`

	// PoolAddress is the address of the associated DAMM v1 pool.
	PoolAddress string `json:"pool_address"`

	// TokenVault is the address of the vault holding the base token deposits.
	TokenVault string `json:"token_vault"`

	// TokenOutVault is the address of the vault holding the output token.
	TokenOutVault string `json:"token_out_vault"`

	// BaseMint is the mint address of the token deposited into the vault.
	BaseMint string `json:"base_mint"`

	// QuoteMint is the mint address of the token received from the vault.
	QuoteMint string `json:"quote_mint"`

	// Base is the base token identifier.
	Base string `json:"base"`

	// Owner is the public key of the vault owner/creator.
	Owner string `json:"owner"`

	// PoolType identifies the pool type (0 = DLMM, 1 = Dynamic AMM, 2 = DAMM v2).
	PoolType int `json:"pool_type"`
}

// ProrataConfig represents the configuration for a prorata alpha vault.
// In prorata mode, all depositors receive tokens proportional to their deposit.
type ProrataConfig struct {
	// Address is the on-chain address of this configuration.
	Address string `json:"address"`

	// MaxBuyingCap is the maximum total deposit amount allowed in native token units.
	MaxBuyingCap int64 `json:"max_buying_cap"`

	// StartVestingDuration is the duration in seconds from activation until vesting begins.
	StartVestingDuration int64 `json:"start_vesting_duration"`

	// EndVestingDuration is the duration in seconds from activation until vesting ends.
	EndVestingDuration int64 `json:"end_vesting_duration"`

	// EscrowFee is the fee charged for creating an escrow account in lamports.
	EscrowFee int `json:"escrow_fee"`

	// ActivationType determines when the vault activates (0 = immediate, 1 = slot-based, 2 = time-based).
	ActivationType int `json:"activation_type"`
}

// FCFSConfig represents the configuration for a first-come-first-served alpha vault.
// In FCFS mode, deposits are accepted until the cap is reached.
type FCFSConfig struct {
	// Address is the on-chain address of this configuration.
	Address string `json:"address"`

	// MaxDepositingCap is the maximum total deposit amount in native token units.
	MaxDepositingCap int64 `json:"max_depositing_cap"`

	// StartVestingDuration is the duration in seconds from activation until vesting begins.
	StartVestingDuration int64 `json:"start_vesting_duration"`

	// EndVestingDuration is the duration in seconds from activation until vesting ends.
	EndVestingDuration int64 `json:"end_vesting_duration"`

	// DepositingDurationUntilLastJoinPoint is the duration in seconds during which
	// deposits are accepted, measured from vault activation.
	DepositingDurationUntilLastJoinPoint int64 `json:"depositing_duration_until_last_join_point"`

	// IndividualDepositingCap is the maximum deposit amount per wallet in native token units.
	IndividualDepositingCap int64 `json:"individual_depositing_cap"`

	// EscrowFee is the fee charged for creating an escrow account in lamports.
	EscrowFee int `json:"escrow_fee"`

	// ActivationType determines when the vault activates (0 = immediate, 1 = slot-based, 2 = time-based).
	ActivationType int `json:"activation_type"`
}

// AlphaVaultConfigs holds both prorata and FCFS alpha vault configurations.
type AlphaVaultConfigs struct {
	// ProrataConfigs contains configurations for prorata-mode alpha vaults.
	ProrataConfigs []ProrataConfig `json:"prorata_configs"`

	// FCFSConfigs contains configurations for FCFS-mode alpha vaults.
	FCFSConfigs []FCFSConfig `json:"fcfs_configs"`
}
