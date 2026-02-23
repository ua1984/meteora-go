// Package dammv1 provides a client for the Meteora DAMM v1 (Dynamic AMM v1) API.
//
// DAMM v1 is Meteora's original constant-product AMM. It supports standard and stable
// swap pools, multitoken pools, farming, and alpha vaults for token launches. Many numeric
// values in the API responses are returned as strings to preserve precision.
//
// The API is served from https://amm-v2.meteora.ag with a rate limit of 10 requests
// per second.
//
// # Endpoints
//
//   - GET /pools: List pools, optionally filtered by address.
//   - GET /pools/search: Search pools with pagination and sorting.
//   - GET /pools-metrics: Protocol-level aggregate metrics.
//   - GET /pool-configs: List all pool configurations.
//   - GET /fee-config/{address}: Fee configuration for a specific config.
//   - GET /farm: List pools with active farming programs.
//   - GET /alpha-vault: List all alpha vaults.
//   - GET /alpha-vault-configs: List alpha vault configurations (prorata and FCFS).
//   - POST /pools_by_a_vault_lp: Find pools by vault LP address.
//
// # Usage
//
// Create a client via the top-level meteora package:
//
//	client := meteora.New()
//	pools, err := client.DAMMv1.ListPools(ctx, "")
//
// Or create a standalone DAMM v1 client:
//
//	httpClient := httpclient.New("https://amm-v2.meteora.ag", nil)
//	client := dammv1.NewClient(httpClient)
package dammv1
