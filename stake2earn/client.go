package stake2earn

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ua1984/meteora-go/internal/httpclient"
)

// Client provides access to the Stake2Earn API.
type Client struct {
	http *httpclient.Client
}

// NewClient creates a new Stake2Earn client.
func NewClient(http *httpclient.Client) *Client {
	return &Client{http: http}
}

// GetAnalytics returns protocol-wide Stake2Earn analytics.
func (c *Client) GetAnalytics(ctx context.Context) (*Analytics, error) {
	var analytics Analytics
	if err := c.http.Get(ctx, "/analytics/all", nil, &analytics); err != nil {
		return nil, fmt.Errorf("stake2earn.GetAnalytics: %w", err)
	}

	return &analytics, nil
}

// ListVaults returns all Stake2Earn vaults.
func (c *Client) ListVaults(ctx context.Context) (*VaultListResponse, error) {
	var resp VaultListResponse
	if err := c.http.Get(ctx, "/vault/all", nil, &resp); err != nil {
		return nil, fmt.Errorf("stake2earn.ListVaults: %w", err)
	}

	return &resp, nil
}

// FilterVaults returns filtered and paginated vaults.
func (c *Client) FilterVaults(ctx context.Context, params *FilterParams) (*VaultListResponse, error) {
	q := url.Values{}
	if params != nil {
		if params.Page != nil {
			q.Set("page", strconv.Itoa(*params.Page))
		}
		if params.Size != nil {
			q.Set("size", strconv.Itoa(*params.Size))
		}
		if params.SortBy != nil {
			q.Set("sort_by", *params.SortBy)
		}
		if params.SortOrder != nil {
			q.Set("sort_order", *params.SortOrder)
		}
		if params.SearchTerm != nil {
			q.Set("search_term", *params.SearchTerm)
		}
	}

	var resp VaultListResponse
	if err := c.http.Get(ctx, "/vault/filter", q, &resp); err != nil {
		return nil, fmt.Errorf("stake2earn.FilterVaults: %w", err)
	}

	return &resp, nil
}

// GetVault returns a single vault by address.
func (c *Client) GetVault(ctx context.Context, address string) (*Vault, error) {
	path := fmt.Sprintf("/vault/%s", address)
	var vault Vault
	if err := c.http.Get(ctx, path, nil, &vault); err != nil {
		return nil, fmt.Errorf("stake2earn.GetVault: %w", err)
	}

	return &vault, nil
}
