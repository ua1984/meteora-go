package dammv2

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ua1984/meteora-go/internal/httpclient"
)

// Client provides access to the DAMM v2 API.
type Client struct {
	http *httpclient.Client
}

// NewClient creates a new DAMM v2 client.
func NewClient(http *httpclient.Client) *Client {
	return &Client{http: http}
}

// ListPools returns a paginated list of DAMM v2 pools.
func (c *Client) ListPools(ctx context.Context, params *ListPoolsParams) (*PaginatedResponse[Pool], error) {
	q := url.Values{}
	if params != nil {
		if params.Page != nil {
			q.Set("page", strconv.Itoa(*params.Page))
		}
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
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
		if params.HideBlacklist != nil {
			q.Set("hide_blacklist", strconv.FormatBool(*params.HideBlacklist))
		}
		if params.IncludeTags != nil {
			q.Set("include_tags", *params.IncludeTags)
		}
		if params.ExcludeTags != nil {
			q.Set("exclude_tags", *params.ExcludeTags)
		}
	}

	var resp PaginatedResponse[Pool]
	if err := c.http.Get(ctx, "/pools", q, &resp); err != nil {
		return nil, fmt.Errorf("dammv2.ListPools: %w", err)
	}

	return &resp, nil
}

// ListGroups returns a paginated list of pool groups.
func (c *Client) ListGroups(ctx context.Context, params *ListGroupsParams) (*PaginatedResponse[PoolGroup], error) {
	q := url.Values{}
	if params != nil {
		if params.Page != nil {
			q.Set("page", strconv.Itoa(*params.Page))
		}
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
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
		if params.HideBlacklist != nil {
			q.Set("hide_blacklist", strconv.FormatBool(*params.HideBlacklist))
		}
	}

	var resp PaginatedResponse[PoolGroup]
	if err := c.http.Get(ctx, "/pools/groups", q, &resp); err != nil {
		return nil, fmt.Errorf("dammv2.ListGroups: %w", err)
	}

	return &resp, nil
}

// GetGroup returns pools within a specific token pair group.
func (c *Client) GetGroup(ctx context.Context, lexicalOrderMints string, params *GetGroupParams) (*PaginatedResponse[Pool], error) {
	q := url.Values{}
	if params != nil {
		if params.Page != nil {
			q.Set("page", strconv.Itoa(*params.Page))
		}
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.SortBy != nil {
			q.Set("sort_by", *params.SortBy)
		}
		if params.SortOrder != nil {
			q.Set("sort_order", *params.SortOrder)
		}
	}

	path := fmt.Sprintf("/pools/groups/%s", lexicalOrderMints)
	var resp PaginatedResponse[Pool]
	if err := c.http.Get(ctx, path, q, &resp); err != nil {
		return nil, fmt.Errorf("dammv2.GetGroup: %w", err)
	}

	return &resp, nil
}

// GetPool returns a single pool by address.
func (c *Client) GetPool(ctx context.Context, address string) (*Pool, error) {
	path := fmt.Sprintf("/pools/%s", address)
	var pool Pool
	if err := c.http.Get(ctx, path, nil, &pool); err != nil {
		return nil, fmt.Errorf("dammv2.GetPool: %w", err)
	}

	return &pool, nil
}

// GetOHLCV returns OHLCV candlestick data for a pool.
func (c *Client) GetOHLCV(ctx context.Context, address string, params *OHLCVParams) (*OHLCVResponse, error) {
	q := url.Values{}
	if params != nil {
		if params.Resolution != nil {
			q.Set("resolution", *params.Resolution)
		}
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
	}

	path := fmt.Sprintf("/pools/%s/ohlcv", address)
	var resp OHLCVResponse
	if err := c.http.Get(ctx, path, q, &resp); err != nil {
		return nil, fmt.Errorf("dammv2.GetOHLCV: %w", err)
	}

	return &resp, nil
}

// GetVolumeHistory returns volume history for a pool.
func (c *Client) GetVolumeHistory(ctx context.Context, address string, params *VolumeHistoryParams) (*VolumeHistoryResponse, error) {
	q := url.Values{}
	if params != nil {
		if params.Resolution != nil {
			q.Set("resolution", *params.Resolution)
		}
		if params.Limit != nil {
			q.Set("limit", strconv.Itoa(*params.Limit))
		}
	}

	path := fmt.Sprintf("/pools/%s/volume/history", address)
	var resp VolumeHistoryResponse
	if err := c.http.Get(ctx, path, q, &resp); err != nil {
		return nil, fmt.Errorf("dammv2.GetVolumeHistory: %w", err)
	}

	return &resp, nil
}

// GetProtocolMetrics returns protocol-wide metrics.
func (c *Client) GetProtocolMetrics(ctx context.Context) (*ProtocolMetrics, error) {
	var metrics ProtocolMetrics
	if err := c.http.Get(ctx, "/stats/protocol_metrics", nil, &metrics); err != nil {
		return nil, fmt.Errorf("dammv2.GetProtocolMetrics: %w", err)
	}

	return &metrics, nil
}
