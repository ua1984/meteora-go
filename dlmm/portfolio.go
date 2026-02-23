package dlmm

// GetPortfolioResponse is the user portfolio response with aggregated pool data.
type GetPortfolioResponse struct {
	HasNext    bool                `json:"hasNext"`
	Page       int                 `json:"page"`
	PageSize   int                 `json:"pageSize"`
	Pools      []PoolPortfolioItem `json:"pools"`
	TotalCount int64               `json:"totalCount"`
}

// PoolPortfolioItem is the aggregated portfolio data for a single pool in the portfolio.
type PoolPortfolioItem struct {
	BaseFee                  string `json:"baseFee"`
	BinStep                  string `json:"binStep"`
	LastClosedAt             *int64 `json:"lastClosedAt"`
	PnLPctChange             string `json:"pnlPctChange"`
	PnLUsd                   string `json:"pnlUsd"`
	PoolAddress              string `json:"poolAddress"`
	TokenX                   string `json:"tokenX"`
	TokenXIcon               string `json:"tokenXIcon"`
	TokenY                   string `json:"tokenY"`
	TokenYIcon               string `json:"tokenYIcon"`
	TotalDeposit             string `json:"totalDeposit"`
	TotalDepositTokenX       string `json:"totalDepositTokenX"`
	TotalDepositTokenXUsd    string `json:"totalDepositTokenXUsd"`
	TotalDepositTokenY       string `json:"totalDepositTokenY"`
	TotalDepositTokenYUsd    string `json:"totalDepositTokenYUsd"`
	TotalFee                 string `json:"totalFee"`
	TotalFeeTokenX           string `json:"totalFeeTokenX"`
	TotalFeeTokenXUsd        string `json:"totalFeeTokenXUsd"`
	TotalFeeTokenY           string `json:"totalFeeTokenY"`
	TotalFeeTokenYUsd        string `json:"totalFeeTokenYUsd"`
	TotalWithdrawal          string `json:"totalWithdrawal"`
	TotalWithdrawalTokenX    string `json:"totalWithdrawalTokenX"`
	TotalWithdrawalTokenXUsd string `json:"totalWithdrawalTokenXUsd"`
	TotalWithdrawalTokenY    string `json:"totalWithdrawalTokenY"`
	TotalWithdrawalTokenYUsd string `json:"totalWithdrawalTokenYUsd"`
}

// GetOpenPortfolioResponse is the user open portfolio response with pool metadata and total metrics.
type GetOpenPortfolioResponse struct {
	HasNext    bool                    `json:"hasNext"`
	Page       int                     `json:"page"`
	PageSize   int                     `json:"pageSize"`
	Pools      []PoolOpenPortfolioItem `json:"pools"`
	SolPrice   *string                 `json:"solPrice"`
	Total      *TotalMetrics           `json:"total"`
	TotalCount int64                   `json:"totalCount"`
}

// PoolOpenPortfolioItem is the aggregated portfolio data for a single pool with open positions.
type PoolOpenPortfolioItem struct {
	Balances                    string   `json:"balances"`
	BalancesSol                 *string  `json:"balancesSol"`
	BaseFee                     float64  `json:"baseFee"`
	BinStep                     int      `json:"binStep"`
	FeePerTVL24h                string   `json:"feePerTvl24h"`
	ListPositions               []string `json:"listPositions"`
	OpenPositionCount           int64    `json:"openPositionCount"`
	OutOfRange                  *bool    `json:"outOfRange"`
	PnL                         string   `json:"pnl"`
	PnLPctChange                string   `json:"pnlPctChange"`
	PnLSol                      *string  `json:"pnlSol"`
	PoolAddress                 string   `json:"poolAddress"`
	PoolPrice                   *float64 `json:"poolPrice"`
	PoolStateUpdatedAtBlockTime *int64   `json:"poolStateUpdatedAtBlockTime"`
	PoolStateUpdatedAtSlot      *int64   `json:"poolStateUpdatedAtSlot"`
	PositionsOutOfRange         []string `json:"positionsOutOfRange"`
	RewardX                     string   `json:"rewardX"`
	RewardY                     string   `json:"rewardY"`
	TokenX                      string   `json:"tokenX"`
	TokenXIcon                  string   `json:"tokenXIcon"`
	TokenXMint                  string   `json:"tokenXMint"`
	TokenY                      string   `json:"tokenY"`
	TokenYIcon                  string   `json:"tokenYIcon"`
}

// TotalMetrics contains total metrics related to user portfolio.
type TotalMetrics struct {
	Balances         string  `json:"balances"`
	BalancesSol      *string `json:"balancesSol"`
	PnL              string  `json:"pnl"`
	PnLSol           *string `json:"pnlSol"`
	UnclaimedFees    string  `json:"unclaimedFees"`
	UnclaimedFeesSol *string `json:"unclaimedFeesSol"`
}

// PortfolioTotalResponse is the portfolio total PnL response.
type PortfolioTotalResponse struct {
	TotalPnLPctChange string `json:"totalPnlPctChange"`
	TotalPnLUsd       string `json:"totalPnlUsd"`
}
