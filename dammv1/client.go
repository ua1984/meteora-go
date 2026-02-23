package dammv1

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ua1984/meteora-go/internal/httpclient"
)

// Client provides access to the DAMM v1 API.
type Client struct {
	http *httpclient.Client
}

// NewClient creates a new DAMM v1 client.
func NewClient(http *httpclient.Client) *Client {
	return &Client{http: http}
}

// ListPools returns pools, optionally filtered by address.
func (c *Client) ListPools(ctx context.Context, address string) ([]Pool, error) {
	q := url.Values{}
	if address != "" {
		q.Set("address", address)
	}

	var pools []Pool
	if err := c.http.Get(ctx, "/pools", q, &pools); err != nil {
		return nil, fmt.Errorf("dammv1.ListPools: %w", err)
	}

	return pools, nil
}

// SearchPools searches for pools with pagination.
func (c *Client) SearchPools(ctx context.Context, params *SearchParams) (*SearchResult, error) {
	q := url.Values{}
	if params != nil {
		if params.Page != nil {
			q.Set("page", strconv.Itoa(*params.Page))
		}
		if params.Size != nil {
			q.Set("size", strconv.Itoa(*params.Size))
		}
		if params.SearchTerm != nil {
			q.Set("search_term", *params.SearchTerm)
		}
		if params.SortBy != nil {
			q.Set("sort_by", *params.SortBy)
		}
		if params.SortOrder != nil {
			q.Set("sort_order", *params.SortOrder)
		}
	}

	var result SearchResult
	if err := c.http.Get(ctx, "/pools/search", q, &result); err != nil {
		return nil, fmt.Errorf("dammv1.SearchPools: %w", err)
	}

	return &result, nil
}

// GetPoolsMetrics returns protocol-level pool metrics.
func (c *Client) GetPoolsMetrics(ctx context.Context) (*PoolMetrics, error) {
	var metrics PoolMetrics
	if err := c.http.Get(ctx, "/pools-metrics", nil, &metrics); err != nil {
		return nil, fmt.Errorf("dammv1.GetPoolsMetrics: %w", err)
	}

	return &metrics, nil
}

// ListPoolConfigs returns all pool configurations.
func (c *Client) ListPoolConfigs(ctx context.Context) ([]PoolConfig, error) {
	var configs []PoolConfig
	if err := c.http.Get(ctx, "/pool-configs", nil, &configs); err != nil {
		return nil, fmt.Errorf("dammv1.ListPoolConfigs: %w", err)
	}

	return configs, nil
}

// GetFeeConfig returns fee configurations for a config address.
func (c *Client) GetFeeConfig(ctx context.Context, configAddr string) ([]FeeConfig, error) {
	path := fmt.Sprintf("/fee-config/%s", configAddr)

	var configs []FeeConfig
	if err := c.http.Get(ctx, path, nil, &configs); err != nil {
		return nil, fmt.Errorf("dammv1.GetFeeConfig: %w", err)
	}

	return configs, nil
}

// ListPoolsWithFarm returns pools that have farming, with pagination.
func (c *Client) ListPoolsWithFarm(ctx context.Context, params *PaginationParams) ([]Pool, error) {
	q := url.Values{}
	if params != nil {
		if params.Page != nil {
			q.Set("page", strconv.Itoa(*params.Page))
		}
		if params.Size != nil {
			q.Set("size", strconv.Itoa(*params.Size))
		}
	}

	var pools []Pool
	if err := c.http.Get(ctx, "/farm", q, &pools); err != nil {
		return nil, fmt.Errorf("dammv1.ListPoolsWithFarm: %w", err)
	}

	return pools, nil
}

// ListAlphaVaults returns alpha vaults, optionally filtered by vault address, pool address, or base mint.
func (c *Client) ListAlphaVaults(ctx context.Context, params *AlphaVaultParams) ([]AlphaVault, error) {
	var q url.Values
	if params != nil {
		q = url.Values{}
		for _, v := range params.VaultAddress {
			q.Add("vault_address", v)
		}
		for _, v := range params.PoolAddress {
			q.Add("pool_address", v)
		}
		for _, v := range params.BaseMint {
			q.Add("base_mint", v)
		}
		if len(q) == 0 {
			q = nil
		}
	}

	var vaults []AlphaVault
	if err := c.http.Get(ctx, "/alpha-vault", q, &vaults); err != nil {
		return nil, fmt.Errorf("dammv1.ListAlphaVaults: %w", err)
	}

	return vaults, nil
}

// ListAlphaVaultConfigs returns all alpha vault configurations.
func (c *Client) ListAlphaVaultConfigs(ctx context.Context) (*AlphaVaultConfigs, error) {
	var configs AlphaVaultConfigs
	if err := c.http.Get(ctx, "/alpha-vault-configs", nil, &configs); err != nil {
		return nil, fmt.Errorf("dammv1.ListAlphaVaultConfigs: %w", err)
	}

	return &configs, nil
}

// GetPoolsByVaultLP returns pools associated with a vault LP address.
func (c *Client) GetPoolsByVaultLP(ctx context.Context, address string) ([]Pool, error) {
	q := url.Values{}
	q.Set("address", address)

	var pools []Pool
	if err := c.http.Post(ctx, "/get_pools_by_a_vault_lp", q, &pools); err != nil {
		return nil, fmt.Errorf("dammv1.GetPoolsByVaultLP: %w", err)
	}

	return pools, nil
}
