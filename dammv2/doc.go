// Package dammv2 provides a client for the Meteora DAMM v2 (Dynamic AMM v2) API.
//
// DAMM v2 pools are the next generation of Meteora's constant-product AMM, supporting
// features like concentrated liquidity ranges, permanent lock liquidity, vested liquidity
// with time-based unlock periods, and alpha vaults for token launches.
//
// The API is served from https://damm-v2.datapi.meteora.ag with a rate limit of
// 10 requests per second. The endpoint structure mirrors the DLMM datapi, providing
// paginated pool listing, pool groups, OHLCV data, volume history, and protocol metrics.
//
// # Usage
//
// Create a client via the top-level meteora package:
//
//	client := meteora.New()
//	pools, err := client.DAMMv2.ListPools(ctx, &dammv2.ListPoolsParams{
//	    Limit: intPtr(10),
//	})
//
// Or create a standalone DAMM v2 client:
//
//	httpClient := httpclient.New("https://damm-v2.datapi.meteora.ag", nil)
//	client := dammv2.NewClient(httpClient)
package dammv2
