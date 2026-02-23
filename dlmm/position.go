package dlmm

// PositionEventType represents the type of a position event.
type PositionEventType string

const (
	PositionEventTypeAdd          PositionEventType = "add"
	PositionEventTypeRemove       PositionEventType = "remove"
	PositionEventTypeClaimFee     PositionEventType = "claim_fee"
	PositionEventTypeClaimReward  PositionEventType = "claim_reward"
)

// PositionEventOrderDirection represents the sort order for position events.
type PositionEventOrderDirection string

const (
	PositionEventOrderDirectionAsc  PositionEventOrderDirection = "asc"
	PositionEventOrderDirectionDesc PositionEventOrderDirection = "desc"
)

// PositionStatus represents the status filter for positions.
type PositionStatus string

const (
	PositionStatusOpen   PositionStatus = "open"
	PositionStatusClosed PositionStatus = "closed"
	PositionStatusAll    PositionStatus = "all"
)

// PositionEvent represents a single historical event for a position.
type PositionEvent struct {
	AmountX         string `json:"amountX"`
	AmountXUsd      string `json:"amountXUsd"`
	AmountY         string `json:"amountY"`
	AmountYUsd      string `json:"amountYUsd"`
	BlockTime       int64  `json:"blockTime"`
	CreatedAt       string `json:"createdAt"`
	EventType       string `json:"eventType"`
	IxIndex         int64  `json:"ixIndex"`
	PoolAddress     string `json:"poolAddress"`
	PositionAddress string `json:"positionAddress"`
	Signature       string `json:"signature"`
	Slot            int64  `json:"slot"`
	TokenX          string `json:"tokenX"`
	TokenY          string `json:"tokenY"`
	TotalUsd        string `json:"totalUsd"`
	UserAddress     string `json:"userAddress"`
}

// GetPositionHistoricalEventsResponse is the response for the GetPositionHistoricalEvents endpoint.
type GetPositionHistoricalEventsResponse struct {
	Events []PositionEvent `json:"events"`
}

// PositionTotalClaimFees represents the total claim fees for a position.
type PositionTotalClaimFees struct {
	ClaimCount    int64  `json:"claimCount"`
	LastClaimTime string `json:"lastClaimTime"`
	PoolAddress   string `json:"poolAddress"`
	TokenX        string `json:"tokenX"`
	TokenY        string `json:"tokenY"`
	TotalFeeX     string `json:"totalFeeX"`
	TotalFeeXUsd  string `json:"totalFeeXUsd"`
	TotalFeeY     string `json:"totalFeeY"`
	TotalFeeYUsd  string `json:"totalFeeYUsd"`
	UserAddress   string `json:"userAddress"`
}

// TokenAmount represents an amount in token units with USD and optional SOL equivalents.
type TokenAmount struct {
	Amount    string  `json:"amount"`
	AmountSol *string `json:"amountSol"`
	Usd       string  `json:"usd"`
}

// TotalUsd represents a total amount in USD with optional SOL equivalent.
type TotalUsd struct {
	Sol *string `json:"sol"`
	Usd string  `json:"usd"`
}

// TokenPairWithTotal represents token X and Y amounts along with a combined total.
type TokenPairWithTotal struct {
	TokenX TokenAmount `json:"tokenX"`
	TokenY TokenAmount `json:"tokenY"`
	Total  TotalUsd    `json:"total"`
}

// UnrealizedPnL represents the live (unrealized) PnL for an open position.
type UnrealizedPnL struct {
	BalanceTokenX         TokenAmount `json:"balanceTokenX"`
	BalanceTokenY         TokenAmount `json:"balanceTokenY"`
	Balances              float64     `json:"balances"`
	BalancesSol           *string     `json:"balancesSol"`
	UnclaimedFeeTokenX    TokenAmount `json:"unclaimedFeeTokenX"`
	UnclaimedFeeTokenY    TokenAmount `json:"unclaimedFeeTokenY"`
	UnclaimedRewardTokenX TokenAmount `json:"unclaimedRewardTokenX"`
	UnclaimedRewardTokenY TokenAmount `json:"unclaimedRewardTokenY"`
}

// PositionPnLData contains PnL data for a single position.
type PositionPnLData struct {
	AllTimeDeposits    TokenPairWithTotal `json:"allTimeDeposits"`
	AllTimeFees        TokenPairWithTotal `json:"allTimeFees"`
	AllTimeWithdrawals TokenPairWithTotal `json:"allTimeWithdrawals"`
	ClosedAt           *int64             `json:"closedAt"`
	CreatedAt          *int64             `json:"createdAt"`
	FeePerTVL24h       string             `json:"feePerTvl24h"`
	IsClosed           bool               `json:"isClosed"`
	IsOutOfRange       *bool              `json:"isOutOfRange"`
	LowerBinId         int32              `json:"lowerBinId"`
	MaxPrice           string             `json:"maxPrice"`
	MinPrice           string             `json:"minPrice"`
	PnLPctChange       string             `json:"pnlPctChange"`
	PnLSol             *string            `json:"pnlSol"`
	PnLUsd             string             `json:"pnlUsd"`
	PoolActiveBinId    *int32             `json:"poolActiveBinId"`
	PoolActivePrice    *string            `json:"poolActivePrice"`
	PositionAddress    string             `json:"positionAddress"`
	UnrealizedPnL      *UnrealizedPnL     `json:"unrealizedPnl"`
	UpperBinId         int32              `json:"upperBinId"`
}

// GetPoolPositionPnLResponse is the response for the GetPoolPositionPnL endpoint.
type GetPoolPositionPnLResponse struct {
	HasNext           bool              `json:"hasNext"`
	Page              int               `json:"page"`
	PageSize          int               `json:"pageSize"`
	Positions         []PositionPnLData `json:"positions"`
	RewardTokenX      *string           `json:"rewardTokenX"`
	RewardTokenXPrice string            `json:"rewardTokenXPrice"`
	RewardTokenY      *string           `json:"rewardTokenY"`
	RewardTokenYPrice string            `json:"rewardTokenYPrice"`
	SolPrice          *string           `json:"solPrice"`
	TokenX            *string           `json:"tokenX"`
	TokenXPrice       string            `json:"tokenXPrice"`
	TokenY            *string           `json:"tokenY"`
	TokenYPrice       string            `json:"tokenYPrice"`
	TotalCount        int64             `json:"totalCount"`
}
