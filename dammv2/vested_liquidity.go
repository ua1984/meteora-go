package dammv2

// VestedLiquidity holds the vested liquidity amounts for a DAMM v2 pool,
// broken down by vesting period. Vested liquidity represents locked LP tokens
// that unlock over time.
type VestedLiquidity struct {
	// Months3 is the amount of liquidity vested with a 3-month unlock period in USD.
	Months3 float64 `json:"months_3"`

	// Months6 is the amount of liquidity vested with a 6-month unlock period in USD.
	Months6 float64 `json:"months_6"`
}
