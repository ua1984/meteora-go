# Meteora Go SDK

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/ua1984/meteora-go/test.yml?branch=master&style=flat-square)](https://github.com/ua1984/meteora-go/workflows/test.yml)
[![GoDoc](https://pkg.go.dev/badge/mod/github.com/ua1984/meteora-go)](https://pkg.go.dev/mod/github.com/ua1984/meteora-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/ua1984/meteora-go)](https://goreportcard.com/report/github.com/ua1984/meteora-go)
[![Release](https://img.shields.io/github/release/ua1984/meteora-go.svg?style=flat-square)](https://github.com/ua1984/meteora-go/releases/latest)

Go client library for the [Meteora](https://meteora.ag) REST APIs. Covers all five Meteora services with a unified client and zero external dependencies (only `net/http`).

## Installation

```bash
go get github.com/ua1984/meteora-go
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	meteora "github.com/ua1984/meteora-go"
	"github.com/ua1984/meteora-go/dlmm"
)

func main() {
	client := meteora.New()
	ctx := context.Background()

	page, limit := 1, 5
	pools, err := client.DLMM.ListPools(ctx, &dlmm.ListPoolsParams{
		Page:  &page,
		Limit: &limit,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range pools.Data {
		fmt.Printf("%s - TVL: $%.2f\n", p.Name, p.TVL)
	}
}
```

See [`examples/basic`](examples/basic/main.go) for a full working example covering all services.

## Services

The unified `meteora.Client` exposes five service clients:

| Field          | Package        | Description                                                                      |
| -------------- | -------------- | -------------------------------------------------------------------------------- |
| `DLMM`         | `dlmm`         | Dynamic Liquidity Market Maker - concentrated liquidity with discrete price bins |
| `DAMMv2`       | `dammv2`       | Dynamic AMM v2 - constant-product AMM with vested liquidity and alpha vaults     |
| `DAMMv1`       | `dammv1`       | Dynamic AMM v1 - original AMM with multitoken pools, farming, and stable swaps   |
| `Stake2Earn`   | `stake2earn`   | Fee vaults where users stake LP tokens to earn trading fees                      |
| `DynamicVault` | `dynamicvault` | Auto-yield vaults allocating tokens across lending protocols                     |

### DLMM

Base URLs: `https://dlmm.datapi.meteora.ag` (30 req/s), `https://dlmm-api.meteora.ag` (30 req/s legacy)

```go
client.DLMM.ListPools(ctx, params)           // Paginated pool listing with search/sort/filter
client.DLMM.ListGroups(ctx, params)           // Pool groups by token pair
client.DLMM.GetGroup(ctx, mints, params)      // Pools in a specific group
client.DLMM.GetPool(ctx, address)             // Single pool by address
client.DLMM.GetOHLCV(ctx, address, params)    // Candlestick data (1m, 5m, 15m, 1h, 4h, 1d)
client.DLMM.GetVolumeHistory(ctx, addr, p)    // Volume history
client.DLMM.GetProtocolMetrics(ctx)           // Protocol-wide stats (TVL, volume, fees)
client.DLMM.ListAllPairs(ctx)                 // Legacy flat pair listing
```

### DAMM v2

Base URL: `https://damm-v2.datapi.meteora.ag` (10 req/s)

```go
client.DAMMv2.ListPools(ctx, params)          // Paginated pool listing
client.DAMMv2.ListGroups(ctx, params)         // Pool groups by token pair
client.DAMMv2.GetGroup(ctx, mints, params)    // Pools in a specific group
client.DAMMv2.GetPool(ctx, address)           // Single pool by address
client.DAMMv2.GetOHLCV(ctx, address, params)  // Candlestick data
client.DAMMv2.GetVolumeHistory(ctx, addr, p)  // Volume history
client.DAMMv2.GetProtocolMetrics(ctx)         // Protocol-wide stats
```

### DAMM v1

Base URL: `https://amm-v2.meteora.ag` (10 req/s)

```go
client.DAMMv1.ListPools(ctx, address)         // All pools, optionally filtered by address
client.DAMMv1.SearchPools(ctx, params)        // Search with pagination and sorting
client.DAMMv1.GetPoolsMetrics(ctx)            // Protocol-level aggregate metrics
client.DAMMv1.ListPoolConfigs(ctx)            // All pool configurations
client.DAMMv1.GetFeeConfig(ctx, configAddr)   // Fee config by address
client.DAMMv1.ListPoolsWithFarm(ctx, params)  // Pools with active farming
client.DAMMv1.ListAlphaVaults(ctx)            // All alpha vaults
client.DAMMv1.ListAlphaVaultConfigs(ctx)      // Alpha vault configs (prorata and FCFS)
client.DAMMv1.GetPoolsByVaultLP(ctx, address)  // Pools by vault LP address
```

### Stake2Earn

Base URL: `https://stake-for-fee-api.meteora.ag`

```go
client.Stake2Earn.GetAnalytics(ctx)           // Protocol-wide stats
client.Stake2Earn.ListVaults(ctx)             // All vaults
client.Stake2Earn.FilterVaults(ctx, params)   // Search/sort/paginate vaults
client.Stake2Earn.GetVault(ctx, address)      // Single vault by address
```

### Dynamic Vault

Base URL: `https://merv2-api.meteora.ag`

```go
client.DynamicVault.ListVaultInfo(ctx)                       // All vaults with state, APY, strategies
client.DynamicVault.ListVaultAddresses(ctx)                  // Vault addresses (vault, LP mint, fee)
client.DynamicVault.GetVaultState(ctx, tokenMint)            // Current vault state
client.DynamicVault.GetAPYState(ctx, tokenMint)              // APY breakdown by strategy
client.DynamicVault.GetAPYByTimeRange(ctx, mint, start, end) // Historical APY
client.DynamicVault.GetVirtualPrice(ctx, mint, strategy)     // Virtual price history
```

## Configuration

```go
// Default client
client := meteora.New()

// Custom HTTP client
client := meteora.New(
	meteora.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}),
)

// Override base URLs
client := meteora.New(
	meteora.WithDLMMBaseURL("https://custom-dlmm.example.com"),
	meteora.WithDAMMv2BaseURL("https://custom-damm.example.com"),
)
```

Available options: `WithHTTPClient`, `WithDLMMBaseURL`, `WithDLMMLegacyBaseURL`, `WithDAMMv2BaseURL`, `WithDAMMv1BaseURL`, `WithStake2EarnBaseURL`, `WithDynamicVaultBaseURL`.

## Error Handling

Non-2xx HTTP responses are returned as `*httpclient.APIError`:

```go
pools, err := client.DLMM.ListPools(ctx, nil)
if err != nil {
	var apiErr *httpclient.APIError
	if errors.As(err, &apiErr) {
		fmt.Printf("HTTP %d: %s\n", apiErr.StatusCode, apiErr.Body)
	}
}
```

## Requirements

- Go 1.18 or later

## License

See LICENSE file for details.
