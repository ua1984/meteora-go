package dammv2

// AmountTotals holds token amounts in X, Y, and USD for a position.
type AmountTotals struct {
	// AmountX is the amount in token X units.
	AmountX float64 `json:"amount_x"`

	// AmountY is the amount in token Y units.
	AmountY float64 `json:"amount_y"`

	// AmountUSD is the equivalent amount in USD.
	AmountUSD float64 `json:"amount_usd"`
}

// ClosedPosition represents a closed liquidity position.
type ClosedPosition struct {
	// PositionAddress is the on-chain address of the position account.
	PositionAddress string `json:"position_address"`

	// PoolAddress is the on-chain address of the pool the position belongs to.
	PoolAddress string `json:"pool_address"`

	// IsClosed indicates whether the position has been closed.
	IsClosed bool `json:"is_closed"`

	// CreatedAt is the Unix timestamp when the position was opened.
	CreatedAt int64 `json:"created_at"`

	// ClosedAt is the Unix timestamp when the position was closed.
	ClosedAt int64 `json:"closed_at"`

	// TotalDeposits is the cumulative amount deposited into the position.
	TotalDeposits AmountTotals `json:"total_deposits"`

	// TotalWithdraws is the cumulative amount withdrawn from the position.
	TotalWithdraws AmountTotals `json:"total_withdraws"`

	// TotalClaimedFees is the cumulative fees claimed from the position.
	TotalClaimedFees AmountTotals `json:"total_claimed_fees"`

	// PnL is the realized profit and loss in USD.
	PnL float64 `json:"pnl"`

	// PnLChangePct is the PnL as a percentage of total deposits.
	PnLChangePct float64 `json:"pnl_change_pct"`
}

// CursorPaginatedResponse is a cursor-based paginated response.
type CursorPaginatedResponse[T any] struct {
	// Limit is the maximum number of items returned.
	Limit int64 `json:"limit"`

	// NextCursor is the cursor to pass for the next page, or nil if there are no more pages.
	NextCursor *string `json:"next_cursor"`

	// Data is the list of items in this page.
	Data []T `json:"data"`
}

// CurrentPosition holds the current on-chain state of an open position.
type CurrentPosition struct {
	// CurrentDeposits is the current value of the deposited liquidity.
	CurrentDeposits AmountTotals `json:"current_deposits"`

	// UnclaimedFees is the fees accrued but not yet claimed.
	UnclaimedFees AmountTotals `json:"unclaimed_fees"`

	// UpdatedAtSlot is the Solana slot number when the position was last updated.
	UpdatedAtSlot int64 `json:"updated_at_slot"`
}

// OpenPosition represents an active (non-closed) liquidity position.
type OpenPosition struct {
	// PositionAddress is the on-chain address of the position account.
	PositionAddress string `json:"position_address"`

	// PoolAddress is the on-chain address of the pool the position belongs to.
	PoolAddress string `json:"pool_address"`

	// CreatedAt is the Unix timestamp when the position was opened.
	CreatedAt int64 `json:"created_at"`

	// TotalDeposits is the cumulative amount deposited into the position.
	TotalDeposits AmountTotals `json:"total_deposits"`

	// TotalWithdraws is the cumulative amount withdrawn from the position.
	TotalWithdraws AmountTotals `json:"total_withdraws"`

	// TotalClaimedFees is the cumulative fees claimed from the position.
	TotalClaimedFees AmountTotals `json:"total_claimed_fees"`

	// CurrentPosition is the current on-chain state of the position.
	CurrentPosition CurrentPosition `json:"current_position"`

	// UnrealizedPnL is the unrealized profit and loss in USD.
	UnrealizedPnL float64 `json:"unrealized_pnl"`

	// UnrealizedPnLChangePct is the unrealized PnL as a percentage of total deposits.
	UnrealizedPnLChangePct float64 `json:"unrealized_pnl_change_pct"`
}

// PositionToken is a simplified token representation used in position responses.
type PositionToken struct {
	// Address is the token's mint address.
	Address string `json:"address"`

	// Symbol is the token's ticker symbol (e.g., "SOL").
	Symbol string `json:"symbol"`

	// Name is the human-readable name of the token.
	Name string `json:"name"`

	// Icon is the URL to the token's icon image.
	Icon string `json:"icon"`
}

// PositionsByPool groups open positions belonging to the same pool.
type PositionsByPool struct {
	// PoolAddress is the on-chain address of the pool.
	PoolAddress string `json:"pool_address"`

	// Name is the human-readable pool name (e.g., "SOL-USDC").
	Name string `json:"name"`

	// TokenX is the first token in the pair.
	TokenX PositionToken `json:"token_x"`

	// TokenY is the second token in the pair.
	TokenY PositionToken `json:"token_y"`

	// FeePct is the pool's fee percentage.
	FeePct float64 `json:"fee_pct"`

	// MinPrice is the minimum price of the pool's price range.
	MinPrice float64 `json:"min_price"`

	// MaxPrice is the maximum price of the pool's price range.
	MaxPrice float64 `json:"max_price"`

	// PoolPrice is the pool's current price, or nil if unavailable.
	PoolPrice *float64 `json:"pool_price"`

	// Positions is the list of open positions in this pool.
	Positions []OpenPosition `json:"positions"`
}

// OpenPositionsResponse is the response returned by GetOpenPositions.
type OpenPositionsResponse struct {
	// TotalPositions is the total number of open positions across all pools.
	TotalPositions int64 `json:"total_positions"`

	// TotalPools is the number of pools containing open positions.
	TotalPools int64 `json:"total_pools"`

	// Data is the list of positions grouped by pool.
	Data []PositionsByPool `json:"data"`
}
