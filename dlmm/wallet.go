package dlmm

// AmountTotals represents token X and Y amounts along with a combined USD total.
type AmountTotals struct {
	AmountUsd float64 `json:"amount_usd"`
	AmountX   float64 `json:"amount_x"`
	AmountY   float64 `json:"amount_y"`
}

// ClosedPosition represents a single closed position for a wallet.
type ClosedPosition struct {
	ClosedAt         int64        `json:"closed_at"`
	CreatedAt        int64        `json:"created_at"`
	LowerBinId       int32        `json:"lower_bin_id"`
	PnL              float64      `json:"pnl"`
	PnLChangePct     float64      `json:"pnl_change_pct"`
	PoolAddress      string       `json:"pool_address"`
	PositionAddress  string       `json:"position_address"`
	TotalClaimedFees AmountTotals `json:"total_claimed_fees"`
	TotalDeposits    AmountTotals `json:"total_deposits"`
	TotalWithdraws   AmountTotals `json:"total_withdraws"`
	UpperBinId       int32        `json:"upper_bin_id"`
}

// ClosedPositionsCursorResponse is the cursor-paginated response for closed positions.
type ClosedPositionsCursorResponse struct {
	Data       []ClosedPosition `json:"data"`
	Limit      int64            `json:"limit"`
	NextCursor *string          `json:"next_cursor"`
}

// PositionToken represents basic SPL token metadata used within position data.
type PositionToken struct {
	Address  string `json:"address"`
	Decimals int    `json:"decimals"`
	Icon     string `json:"icon"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
}

// CurrentPosition represents the current state of an open position.
type CurrentPosition struct {
	CurrentDeposits AmountTotals `json:"current_deposits"`
	UnclaimedFees   AmountTotals `json:"unclaimed_fees"`
	UpdatedAtSlot   int64        `json:"updated_at_slot"`
}

// OpenPosition represents a single open position within a pool.
type OpenPosition struct {
	CreatedAt              int64           `json:"created_at"`
	CurrentPosition        CurrentPosition `json:"current_position"`
	LowerBinId             int32           `json:"lower_bin_id"`
	LowerBinPrice          float64         `json:"lower_bin_price"`
	PoolAddress            string          `json:"pool_address"`
	PositionAddress        string          `json:"position_address"`
	TotalClaimedFees       AmountTotals    `json:"total_claimed_fees"`
	TotalDeposits          AmountTotals    `json:"total_deposits"`
	TotalWithdraws         AmountTotals    `json:"total_withdraws"`
	UnrealizedPnLValue     float64         `json:"unrealized_pnl"`
	UnrealizedPnLChangePct float64         `json:"unrealized_pnl_change_pct"`
	UpperBinId             int32           `json:"upper_bin_id"`
	UpperBinPrice          float64         `json:"upper_bin_price"`
}

// PositionsByPool groups open positions by their pool.
type PositionsByPool struct {
	ActiveBinId int32          `json:"active_bin_id"`
	BinStep     int            `json:"bin_step"`
	FeePct      float64        `json:"fee_pct"`
	Name        string         `json:"name"`
	PoolAddress string         `json:"pool_address"`
	PoolPrice   *float64       `json:"pool_price"`
	Positions   []OpenPosition `json:"positions"`
	TokenX      PositionToken  `json:"token_x"`
	TokenY      PositionToken  `json:"token_y"`
}

// OpenPositionsResponse is the response for the GetOpenPositions endpoint.
type OpenPositionsResponse struct {
	Data           []PositionsByPool `json:"data"`
	TotalPools     int64             `json:"total_pools"`
	TotalPositions int64             `json:"total_positions"`
}
