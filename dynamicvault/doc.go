// Package dynamicvault provides a client for the Meteora Dynamic Vault API.
//
// Dynamic Vaults automatically allocate deposited tokens across multiple lending
// protocols (Solend, Marginfi, Kamino, etc.) to optimize yield. Users deposit tokens
// and receive LP tokens representing their share of the vault.
//
// The API is served from https://merv2-api.meteora.ag with no published rate limit.
// All endpoints use path-based parameters rather than query parameters.
//
// # Endpoints
//
//   - GET /vault_info: List all vaults with current state, APY, and strategies.
//   - GET /vault_addresses: List vault addresses (vault, LP mint, fee account).
//   - GET /vault_state/{token_mint}: Current state for a specific vault.
//   - GET /apy_state/{token_mint}: APY breakdown by strategy for a vault.
//   - GET /apy_filter/{token_mint}/{start}/{end}: Historical APY within a time range.
//   - GET /virtual_price/{token_mint}/{strategy}: Virtual price history for a strategy.
//
// # Usage
//
// Create a client via the top-level meteora package:
//
//	client := meteora.New()
//	vaults, err := client.DynamicVault.ListVaultInfo(ctx)
//
// Or create a standalone Dynamic Vault client:
//
//	httpClient := httpclient.New("https://merv2-api.meteora.ag", nil)
//	client := dynamicvault.NewClient(httpClient)
package dynamicvault
