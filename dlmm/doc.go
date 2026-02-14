// Package dlmm provides a client for the Meteora DLMM (Dynamic Liquidity Market Maker) API.
//
// DLMM pools use a concentrated liquidity bin-based model where liquidity is distributed
// across discrete price bins. This enables zero-slippage trades within active bins and
// more capital-efficient liquidity provision compared to traditional AMMs.
//
// This package provides access to two API surfaces:
//
//   - Datapi (https://dlmm.datapi.meteora.ag): The primary API with paginated pool listing,
//     pool groups, OHLCV candlestick data, volume history, and protocol metrics.
//     Rate limit: 30 requests per second.
//
//   - Legacy API (https://dlmm-api.meteora.ag): The legacy /pair/all endpoint that returns
//     all pairs in a flat format. Rate limit: 30 requests per second.
//
// # Usage
//
// Create a client via the top-level meteora package:
//
//	client := meteora.New()
//	pools, err := client.DLMM.ListPools(ctx, &dlmm.ListPoolsParams{
//	    Limit: intPtr(10),
//	})
//
// Or create a standalone DLMM client:
//
//	httpClient := httpclient.New("https://dlmm.datapi.meteora.ag", nil)
//	legacyClient := httpclient.New("https://dlmm-api.meteora.ag", nil)
//	client := dlmm.NewClient(httpClient, legacyClient)
package dlmm
