package stake2earn

// VaultFlags holds flag indicators for a Stake2Earn vault, used to signal
// threshold-based conditions.
type VaultFlags struct {
	// TVLUSDThresholdReached indicates whether the vault's TVL has reached the minimum threshold.
	TVLUSDThresholdReached bool `json:"tvl_usd_threshold_reached"`

	// TVLUSDThreshold is the minimum TVL threshold in USD required for the vault to be active.
	TVLUSDThreshold float64 `json:"tvl_usd_threshold"`
}

// Vault represents a Stake2Earn vault where users stake LP tokens to earn
// a share of the pool's trading fees.
type Vault struct {
	// VaultAddress is the on-chain address of the Stake2Earn vault.
	VaultAddress string `json:"vault_address"`

	// PoolAddress is the address of the underlying trading pool.
	PoolAddress string `json:"pool_address"`

	// TokenAMint is the mint address of the first token in the pair.
	TokenAMint string `json:"token_a_mint"`

	// TokenBMint is the mint address of the second token in the pair.
	TokenBMint string `json:"token_b_mint"`

	// StakeMint is the mint address of the token that can be staked.
	StakeMint string `json:"stake_mint"`

	// TokenASymbol is the ticker symbol of token A (e.g., "SOL").
	TokenASymbol string `json:"token_a_symbol"`

	// TokenBSymbol is the ticker symbol of token B (e.g., "USDC").
	TokenBSymbol string `json:"token_b_symbol"`

	// TotalStakedAmount is the total amount of tokens staked in native units.
	TotalStakedAmount float64 `json:"total_staked_amount"`

	// TotalStakedAmountUSD is the total value of staked tokens in USD.
	TotalStakedAmountUSD float64 `json:"total_staked_amount_usd"`

	// CurrentRewardTokenAUSD is the current pending reward for token A in USD.
	CurrentRewardTokenAUSD float64 `json:"current_reward_token_a_usd"`

	// CurrentRewardTokenBUSD is the current pending reward for token B in USD.
	CurrentRewardTokenBUSD float64 `json:"current_reward_token_b_usd"`

	// CurrentRewardUSD is the total current pending rewards in USD.
	CurrentRewardUSD float64 `json:"current_reward_usd"`

	// DailyRewardUSD is the estimated daily reward distribution in USD.
	DailyRewardUSD float64 `json:"daily_reward_usd"`

	// CreatedAtSlot is the Solana slot number when the vault was created.
	CreatedAtSlot int64 `json:"created_at_slot"`

	// CreatedAtSlotTimestamp is the Unix timestamp of the creation slot.
	CreatedAtSlotTimestamp int64 `json:"created_at_slot_timestamp"`

	// CreatedAtTxSig is the transaction signature of the vault creation transaction.
	CreatedAtTxSig string `json:"created_at_tx_sig"`

	// SecondsToFullUnlock is the number of seconds until all staked tokens can be withdrawn.
	SecondsToFullUnlock int64 `json:"seconds_to_full_unlock"`

	// StartFeeDistributeTimestamp is the Unix timestamp when fee distribution begins.
	StartFeeDistributeTimestamp int64 `json:"start_fee_distribute_timestamp"`

	// MarketCap is the market capitalization of the staked token in USD.
	MarketCap float64 `json:"marketcap"`

	// Flags holds threshold-based condition indicators for the vault.
	Flags VaultFlags `json:"flags"`
}

// VaultListResponse wraps a list of vaults with a total count.
type VaultListResponse struct {
	// Total is the total number of vaults matching the query.
	Total int `json:"total"`

	// Data contains the vault records.
	Data []Vault `json:"data"`
}
