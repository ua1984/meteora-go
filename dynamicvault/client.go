package dynamicvault

import (
	"context"
	"fmt"

	"github.com/ua1984/meteora-go/internal/httpclient"
)

// Client provides access to the Dynamic Vault API.
type Client struct {
	http *httpclient.Client
}

// NewClient creates a new Dynamic Vault client.
func NewClient(http *httpclient.Client) *Client {
	return &Client{http: http}
}

// ListVaultInfo returns information for all vaults.
func (c *Client) ListVaultInfo(ctx context.Context) ([]VaultInfo, error) {
	var vaults []VaultInfo
	if err := c.http.Get(ctx, "/vault_info", nil, &vaults); err != nil {
		return nil, fmt.Errorf("dynamicvault.ListVaultInfo: %w", err)
	}

	return vaults, nil
}

// ListVaultAddresses returns addresses for all vaults.
func (c *Client) ListVaultAddresses(ctx context.Context) ([]VaultAddress, error) {
	var addresses []VaultAddress
	if err := c.http.Get(ctx, "/vault_addresses", nil, &addresses); err != nil {
		return nil, fmt.Errorf("dynamicvault.ListVaultAddresses: %w", err)
	}

	return addresses, nil
}

// GetVaultState returns the current state for a vault identified by token mint.
func (c *Client) GetVaultState(ctx context.Context, tokenMint string) (*VaultState, error) {
	path := fmt.Sprintf("/vault_state/%s", tokenMint)
	var state VaultState
	if err := c.http.Get(ctx, path, nil, &state); err != nil {
		return nil, fmt.Errorf("dynamicvault.GetVaultState: %w", err)
	}

	return &state, nil
}

// GetAPYState returns the current APY state for a vault identified by token mint.
func (c *Client) GetAPYState(ctx context.Context, tokenMint string) (*APYState, error) {
	path := fmt.Sprintf("/apy_state/%s", tokenMint)
	var state APYState
	if err := c.http.Get(ctx, path, nil, &state); err != nil {
		return nil, fmt.Errorf("dynamicvault.GetAPYState: %w", err)
	}

	return &state, nil
}

// GetAPYByTimeRange returns APY entries within a time range for a vault.
func (c *Client) GetAPYByTimeRange(ctx context.Context, tokenMint string, start, end int64) ([]APYEntry, error) {
	path := fmt.Sprintf("/apy_filter/%s/%d/%d", tokenMint, start, end)
	var entries []APYEntry
	if err := c.http.Get(ctx, path, nil, &entries); err != nil {
		return nil, fmt.Errorf("dynamicvault.GetAPYByTimeRange: %w", err)
	}

	return entries, nil
}

// GetVirtualPrice returns virtual price data for a vault and strategy.
func (c *Client) GetVirtualPrice(ctx context.Context, tokenMint string, strategy string) ([]VirtualPrice, error) {
	path := fmt.Sprintf("/virtual_price/%s/%s", tokenMint, strategy)
	var prices []VirtualPrice
	if err := c.http.Get(ctx, path, nil, &prices); err != nil {
		return nil, fmt.Errorf("dynamicvault.GetVirtualPrice: %w", err)
	}

	return prices, nil
}
