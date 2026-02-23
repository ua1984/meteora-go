package dlmm

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ua1984/meteora-go/internal/httpclient"
)

// Client provides access to the DLMM API (both datapi and legacy).
type Client struct {
	http   *httpclient.Client
	legacy *httpclient.Client
}

// NewClient creates a new DLMM client.
func NewClient(http, legacy *httpclient.Client) *Client {
	return &Client{http: http, legacy: legacy}
}

// ListPools returns a paginated list of DLMM pools.
func (c *Client) ListPools(ctx context.Context, params *ListPoolsParams) (*PaginatedResponse[Pool], error) {
	q := url.Values{}
	if params != nil {
		if params.Page != nil {
			q.Set("page", strconv.Itoa(*params.Page))
		}
		if params.PageSize != nil {
			q.Set("page_size", strconv.Itoa(*params.PageSize))
		}
		if params.Query != nil {
			q.Set("query", *params.Query)
		}
		if params.FilterBy != nil {
			q.Set("filter_by", *params.FilterBy)
		}
		if params.SortBy != nil {
			q.Set("sort_by", *params.SortBy)
		}
	}

	var resp PaginatedResponse[Pool]
	if err := c.http.Get(ctx, "/pools", q, &resp); err != nil {
		return nil, fmt.Errorf("dlmm.ListPools: %w", err)
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
		if params.PageSize != nil {
			q.Set("page_size", strconv.Itoa(*params.PageSize))
		}
		if params.Query != nil {
			q.Set("query", *params.Query)
		}
		if params.FilterBy != nil {
			q.Set("filter_by", *params.FilterBy)
		}
		if params.SortBy != nil {
			q.Set("sort_by", *params.SortBy)
		}
		if params.VolumeTW != nil {
			q.Set("volume_tw", *params.VolumeTW)
		}
		if params.FeeTVLRatioTW != nil {
			q.Set("fee_tvl_ratio_tw", *params.FeeTVLRatioTW)
		}
	}

	var resp PaginatedResponse[PoolGroup]
	if err := c.http.Get(ctx, "/pools/groups", q, &resp); err != nil {
		return nil, fmt.Errorf("dlmm.ListGroups: %w", err)
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
		if params.PageSize != nil {
			q.Set("page_size", strconv.Itoa(*params.PageSize))
		}
		if params.Query != nil {
			q.Set("query", *params.Query)
		}
		if params.FilterBy != nil {
			q.Set("filter_by", *params.FilterBy)
		}
		if params.SortBy != nil {
			q.Set("sort_by", *params.SortBy)
		}
	}

	path := fmt.Sprintf("/pools/groups/%s", lexicalOrderMints)
	var resp PaginatedResponse[Pool]
	if err := c.http.Get(ctx, path, q, &resp); err != nil {
		return nil, fmt.Errorf("dlmm.GetGroup: %w", err)
	}

	return &resp, nil
}

// GetPool returns a single pool by address.
func (c *Client) GetPool(ctx context.Context, address string) (*Pool, error) {
	path := fmt.Sprintf("/pools/%s", address)
	var pool Pool
	if err := c.http.Get(ctx, path, nil, &pool); err != nil {
		return nil, fmt.Errorf("dlmm.GetPool: %w", err)
	}

	return &pool, nil
}

// GetOHLCV returns OHLCV candlestick data for a pool.
func (c *Client) GetOHLCV(ctx context.Context, address string, params *OHLCVParams) (*OHLCVResponse, error) {
	q := url.Values{}
	if params != nil {
		if params.Timeframe != nil {
			q.Set("timeframe", *params.Timeframe)
		}
		if params.StartTime != nil {
			q.Set("start_time", strconv.FormatInt(*params.StartTime, 10))
		}
		if params.EndTime != nil {
			q.Set("end_time", strconv.FormatInt(*params.EndTime, 10))
		}
	}

	path := fmt.Sprintf("/pools/%s/ohlcv", address)
	var resp OHLCVResponse
	if err := c.http.Get(ctx, path, q, &resp); err != nil {
		return nil, fmt.Errorf("dlmm.GetOHLCV: %w", err)
	}

	return &resp, nil
}

// GetVolumeHistory returns volume history for a pool.
func (c *Client) GetVolumeHistory(ctx context.Context, address string, params *VolumeHistoryParams) (*VolumeHistoryResponse, error) {
	q := url.Values{}
	if params != nil {
		if params.Timeframe != nil {
			q.Set("timeframe", *params.Timeframe)
		}
		if params.StartTime != nil {
			q.Set("start_time", strconv.FormatInt(*params.StartTime, 10))
		}
		if params.EndTime != nil {
			q.Set("end_time", strconv.FormatInt(*params.EndTime, 10))
		}
	}

	path := fmt.Sprintf("/pools/%s/volume/history", address)
	var resp VolumeHistoryResponse
	if err := c.http.Get(ctx, path, q, &resp); err != nil {
		return nil, fmt.Errorf("dlmm.GetVolumeHistory: %w", err)
	}

	return &resp, nil
}

// GetProtocolMetrics returns protocol-wide metrics.
func (c *Client) GetProtocolMetrics(ctx context.Context) (*ProtocolMetrics, error) {
	var metrics ProtocolMetrics
	if err := c.http.Get(ctx, "/stats/protocol_metrics", nil, &metrics); err != nil {
		return nil, fmt.Errorf("dlmm.GetProtocolMetrics: %w", err)
	}

	return &metrics, nil
}

// ListAllPairs returns all pairs from the legacy DLMM API.
func (c *Client) ListAllPairs(ctx context.Context) ([]LegacyPair, error) {
	var pairs []LegacyPair
	if err := c.legacy.Get(ctx, "/pair/all", nil, &pairs); err != nil {
		return nil, fmt.Errorf("dlmm.ListAllPairs: %w", err)
	}

	return pairs, nil
}
