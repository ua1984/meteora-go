package stake2earn

// Analytics holds protocol-wide aggregate statistics for the Stake2Earn service.
// Retrieved from the /analytics/all endpoint.
type Analytics struct {
	// TotalFeeVaults is the total number of Stake2Earn fee vaults.
	TotalFeeVaults int `json:"total_fee_vaults"`

	// TotalStakedAmountUSD is the total value staked across all vaults in USD.
	TotalStakedAmountUSD float64 `json:"total_staked_amount_usd"`
}
