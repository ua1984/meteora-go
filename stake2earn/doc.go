// Package stake2earn provides a client for the Meteora Stake2Earn API.
//
// Stake2Earn (also known as Stake for Fee) allows users to stake LP tokens
// in fee vaults to earn a share of the underlying pool's trading fees. The
// fees are distributed proportionally to stakers based on their share of the vault.
//
// The API is served from https://stake-for-fee-api.meteora.ag with no published
// rate limit.
//
// # Endpoints
//
//   - GET /analytics/all: Protocol-wide analytics (total vaults, total staked amount).
//   - GET /vault/all: List all Stake2Earn vaults.
//   - GET /vault/filter: Search and filter vaults with pagination.
//   - GET /vault/{address}: Get details for a specific vault.
//
// # Usage
//
// Create a client via the top-level meteora package:
//
//	client := meteora.New()
//	analytics, err := client.Stake2Earn.GetAnalytics(ctx)
//
// Or create a standalone Stake2Earn client:
//
//	httpClient := httpclient.New("https://stake-for-fee-api.meteora.ag", nil)
//	client := stake2earn.NewClient(httpClient)
package stake2earn
