package dlmm

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ua1984/meteora-go/internal/httpclient"
)

// Client provides access to the DLMM API.
type Client struct {
	http *httpclient.Client
}

// NewClient creates a new DLMM client.
func NewClient(http *httpclient.Client) *Client {
	return &Client{http: http}
}

// ListPools returns a paginated list of pools.
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

// GetGroup returns a paginated list of pools that belong to a specific pool group.
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

// GetPool returns metadata and current state for a single pool.
func (c *Client) GetPool(ctx context.Context, address string) (*Pool, error) {
	path := fmt.Sprintf("/pools/%s", address)
	var pool Pool
	if err := c.http.Get(ctx, path, nil, &pool); err != nil {
		return nil, fmt.Errorf("dlmm.GetPool: %w", err)
	}

	return &pool, nil
}

// GetOHLCV returns OHLCV candles for a single pool over a time range.
//
// Notes:
//   - If both start_time and end_time are provided, candles are returned in the range [start_time, end_time].
//   - If only one of start_time or end_time is provided, the missing bound is inferred using the selected timeframe.
//   - If neither is provided, a default range is used based on timeframe.
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

// GetVolumeHistory returns historical volume for a pool aggregated into time buckets.
//
// Notes:
//   - If both start_time and end_time are provided, the result covers the range [start_time, end_time].
//   - If only one of start_time or end_time is provided, the missing bound is inferred using the selected timeframe.
//   - If neither is provided, a default range is used based on timeframe.
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

// GetProtocolMetrics returns aggregated protocol-level metrics across all pools.
func (c *Client) GetProtocolMetrics(ctx context.Context) (*ProtocolMetrics, error) {
	var metrics ProtocolMetrics
	if err := c.http.Get(ctx, "/stats/protocol_metrics", nil, &metrics); err != nil {
		return nil, fmt.Errorf("dlmm.GetProtocolMetrics: %w", err)
	}

	return &metrics, nil
}

// GetPortfolio returns the user's portfolio with pool metadata and aggregated PnL.
func (c *Client) GetPortfolio(ctx context.Context, params *GetPortfolioParams) (*GetPortfolioResponse, error) {
	q := url.Values{}
	if params != nil {
		q.Set("user", params.User)
		if params.Page != nil {
			q.Set("page", strconv.Itoa(*params.Page))
		}
		if params.PageSize != nil {
			q.Set("page_size", strconv.Itoa(*params.PageSize))
		}
		if params.DaysBack != nil {
			q.Set("days_back", strconv.Itoa(*params.DaysBack))
		}
	}

	var resp GetPortfolioResponse
	if err := c.http.Get(ctx, "/portfolio", q, &resp); err != nil {
		return nil, fmt.Errorf("dlmm.GetPortfolio: %w", err)
	}

	return &resp, nil
}

// GetOpenPortfolio returns the user's open portfolio with pool metadata, balances, and total metrics.
func (c *Client) GetOpenPortfolio(ctx context.Context, params *GetOpenPortfolioParams) (*GetOpenPortfolioResponse, error) {
	q := url.Values{}
	if params != nil {
		q.Set("user", params.User)
		if params.Page != nil {
			q.Set("page", strconv.Itoa(*params.Page))
		}
		if params.PageSize != nil {
			q.Set("page_size", strconv.Itoa(*params.PageSize))
		}
		if params.SortDirection != nil {
			q.Set("sort_direction", string(*params.SortDirection))
		}
		if params.SortBy != nil {
			q.Set("sort_by", string(*params.SortBy))
		}
	}

	var resp GetOpenPortfolioResponse
	if err := c.http.Get(ctx, "/portfolio/open", q, &resp); err != nil {
		return nil, fmt.Errorf("dlmm.GetOpenPortfolio: %w", err)
	}

	return &resp, nil
}

// GetPortfolioTotal returns the all-time total PnL across all user's pools.
func (c *Client) GetPortfolioTotal(ctx context.Context, user string) (*PortfolioTotalResponse, error) {
	q := url.Values{}
	q.Set("user", user)

	var resp PortfolioTotalResponse
	if err := c.http.Get(ctx, "/portfolio/total", q, &resp); err != nil {
		return nil, fmt.Errorf("dlmm.GetPortfolioTotal: %w", err)
	}

	return &resp, nil
}
