package dammv2

// Pool represents a DAMM v2 (Dynamic AMM v2) pool. DAMM v2 pools support
// features like concentrated liquidity, permanent lock liquidity, vested
// liquidity, and alpha vaults.
type Pool struct {
	// Address is the on-chain address of the pool account.
	Address string `json:"address"`

	// Name is the human-readable pool name (e.g., "SOL-USDC").
	Name string `json:"name"`

	// TokenX is the metadata for the first token in the pair.
	TokenX Token `json:"token_x"`

	// TokenY is the metadata for the second token in the pair.
	TokenY Token `json:"token_y"`

	// TokenXAmount is the amount of token X in the pool.
	TokenXAmount float64 `json:"token_x_amount"`

	// TokenYAmount is the amount of token Y in the pool.
	TokenYAmount float64 `json:"token_y_amount"`

	// CreatedAt is the Unix timestamp when the pool was created.
	CreatedAt int64 `json:"created_at"`

	// VaultX is the vault address for token X.
	VaultX string `json:"vault_x"`

	// VaultY is the vault address for token Y.
	VaultY string `json:"vault_y"`

	// AlphaVault is the associated alpha vault address, if any.
	AlphaVault string `json:"alpha_vault"`

	// PoolConfig holds the pool's configuration parameters.
	PoolConfig PoolConfig `json:"pool_config"`

	// TVL is the total value locked in the pool in USD.
	TVL float64 `json:"tvl"`

	// CurrentPrice is the current price of token X in terms of token Y.
	CurrentPrice float64 `json:"current_price"`

	// HasFarm indicates whether the pool has an active farming program.
	HasFarm bool `json:"has_farm"`

	// FarmAPR is the annual percentage rate from farming rewards.
	FarmAPR float64 `json:"farm_apr"`

	// FarmAPY is the annual percentage yield from farming rewards (compounded).
	FarmAPY float64 `json:"farm_apy"`

	// PermanentLockLiquidity is the amount of permanently locked liquidity in USD.
	PermanentLockLiquidity float64 `json:"permanent_lock_liquidity"`

	// VestedLiquidity holds vested liquidity amounts by unlock period.
	VestedLiquidity VestedLiquidity `json:"vested_liquidity"`

	// Volume contains trading volume aggregated across time windows.
	Volume TimeBuckets `json:"volume"`

	// Fees contains trading fees aggregated across time windows.
	Fees TimeBuckets `json:"fees"`

	// ProtocolFees contains protocol fees aggregated across time windows.
	ProtocolFees TimeBuckets `json:"protocol_fees"`

	// FeeTVLRatio contains fee-to-TVL ratios across time windows.
	FeeTVLRatio TimeBuckets `json:"fee_tvl_ratio"`

	// IsBlacklisted indicates whether the pool has been flagged/blacklisted.
	IsBlacklisted bool `json:"is_blacklisted"`

	// Launchpad identifies the launchpad that created this pool, if any.
	Launchpad string `json:"launchpad"`

	// Tags is a list of labels associated with the pool.
	Tags []string `json:"tags"`
}
