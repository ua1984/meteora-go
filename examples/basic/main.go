package main

import (
	"context"
	"fmt"
	"log"
	"time"

	meteora "github.com/ua1984/meteora-go"
	"github.com/ua1984/meteora-go/dlmm"
)

func main() {
	client := meteora.New()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// List DLMM pools
	pools, err := client.DLMM.ListPools(ctx, &dlmm.ListPoolsParams{
		Page:     meteora.Int(1),
		PageSize: meteora.Int(5),
	})
	if err != nil {
		log.Fatalf("ListPools: %v", err)
	}

	fmt.Printf("DLMM Pools (total: %d, showing %d):\n", pools.Total, len(pools.Data))
	for _, p := range pools.Data {
		fmt.Printf("  %s (%s) - TVL: $%.2f, 24h Vol: $%.2f\n",
			p.Name, p.Address, p.TVL, p.Volume.Hour24)
	}

	// DAMM v2 pools
	fmt.Println()
	dammPools, err := client.DAMMv2.ListPools(ctx, nil)
	if err != nil {
		log.Fatalf("DAMMv2.ListPools: %v", err)
	}

	fmt.Printf("DAMM v2 Pools (total: %d)\n", dammPools.Total)

	// DAMM v1 metrics
	fmt.Println()
	metrics, err := client.DAMMv1.GetPoolsMetrics(ctx)
	if err != nil {
		log.Fatalf("DAMMv1.GetPoolsMetrics: %v", err)
	}

	fmt.Printf("DAMM v1 Metrics:\n")
	fmt.Printf("  Dynamic AMM TVL: $%.2f\n", metrics.DynamicAMMTVL)
	fmt.Printf("  Dynamic AMM Total Volume: $%.2f\n", metrics.DynamicAMMTotalVolume)

	// Stake2Earn analytics
	fmt.Println()
	analytics, err := client.Stake2Earn.GetAnalytics(ctx)
	if err != nil {
		log.Fatalf("Stake2Earn.GetAnalytics: %v", err)
	}

	fmt.Printf("Stake2Earn: %d vaults, $%.2f total staked\n",
		analytics.TotalFeeVaults, analytics.TotalStakedAmountUSD)

	// Dynamic Vault info
	fmt.Println()
	vaults, err := client.DynamicVault.ListVaultInfo(ctx)
	if err != nil {
		log.Fatalf("DynamicVault.ListVaultInfo: %v", err)
	}

	fmt.Printf("Dynamic Vaults: %d total\n", len(vaults))
	for i, v := range vaults {
		if i >= 5 {
			fmt.Printf("  ... and %d more\n", len(vaults)-5)
			break
		}
		fmt.Printf("  %s - Virtual Price: %s, Closest APY: %.2f%%\n",
			v.Symbol, v.VirtualPrice, v.ClosestAPY)
	}
}
