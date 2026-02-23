package dlmm

// ListPoolsParams are optional query parameters for the ListPools method.
type ListPoolsParams struct {
	// Page is the page number (1-based).
	Page *int `json:"page,omitempty"`

	// PageSize is the number of pools to return per page. Max 1000.
	PageSize *int `json:"page_size,omitempty"`

	// Query is the search query used to match pools by name, tokens, or address.
	Query *string `json:"query,omitempty"`

	// FilterBy specifies the conditions to filter documents by field values.
	//
	// Format: <expr> [&& <expr> ...]
	// Where each expression is: <field><op><value>
	//
	// Allowed fields:
	// - Numeric: tvl, volume_*, fee_*, fee_tvl_ratio_*, apr_*
	// - Boolean: is_blacklisted
	// - Text: pool_address, name, token_x, token_y
	//
	// Operators:
	// - Numeric: =, >, >=, <, <=
	// - Boolean: =true, =false
	// - Text:
	//   - exact match: =<value>
	//   - multi-value OR: =[value1|value2|...]
	//
	// Examples:
	// - tvl>1000
	// - is_blacklisted=false && volume_24h>=50000
	FilterBy *string `json:"filter_by,omitempty"`

	// SortBy is used to sort results by one or more fields.
	//
	// Format:
	// - Time-windowed metrics: <metric>_<window>:<direction>
	// - Non-windowed metrics: <field>:<direction>
	//
	// - direction: asc or desc
	// - window (when applicable): 5m, 30m, 1h, 2h, 4h, 12h, 24h
	//
	// Available fields:
	// - Time-windowed metrics: volume_*, fee_*, fee_tvl_ratio_*, apr_*
	// - Non-windowed metrics: tvl, fee_pct, bin_step, pool_created_at, farm_apy
	//
	// Default: volume_24h:desc
	//
	// Examples:
	// - volume_24h:desc
	// - fee_1h:asc
	// - tvl:desc
	SortBy *string `json:"sort_by,omitempty"`
}

// ListGroupsParams are optional query parameters for the ListGroups method.
type ListGroupsParams struct {
	// Page is the page number (1-based).
	Page *int `json:"page,omitempty"`

	// PageSize is the number of pools to return per page. Max 100.
	PageSize *int `json:"page_size,omitempty"`

	// Query is the search query used to match pools by name, tokens, or address.
	Query *string `json:"query,omitempty"`

	// FilterBy specifies the conditions to filter documents by field values.
	//
	// Format: <expr> [&& <expr> ...]
	// Where each expression is: <field><op><value>
	//
	// Allowed fields:
	// - Numeric: tvl, volume_*, fee_*, fee_tvl_ratio_*, apr_*
	// - Boolean: is_blacklisted
	// - Text: pool_address, name, token_x, token_y
	//
	// Operators:
	// - Numeric: =, >, >=, <, <=
	// - Boolean: =true, =false
	// - Text:
	//   - exact match: =<value>
	//   - multi-value OR: =[value1|value2|...]
	//
	// Examples:
	// - tvl>1000
	// - is_blacklisted=false && volume_24h>=50000
	FilterBy *string `json:"filter_by,omitempty"`

	// SortBy is used to sort results by one or more fields.
	//
	// Format:
	// - Time-windowed metrics: <metric>_<window>:<direction>
	// - Non-windowed metrics: <field>:<direction>
	//
	// - direction: asc or desc
	// - window (when applicable): 5m, 30m, 1h, 2h, 4h, 12h, 24h
	//
	// Available fields:
	// - Time-windowed metrics: volume_*, fee_*, fee_tvl_ratio_*, apr_*
	// - Non-windowed metrics: tvl, fee_pct, bin_step, pool_created_at, farm_apy
	//
	// Default: volume_24h:desc
	//
	// Examples:
	// - volume_24h:desc
	// - fee_1h:asc
	// - tvl:desc
	SortBy *string `json:"sort_by,omitempty"`

	// VolumeTW is the time window to aggregate volume. Returns sum. Default: volume_24h.
	VolumeTW *string `json:"volume_tw,omitempty"`

	// FeeTVLRatioTW is the time window to aggregate fee TVL ratio. Returns Max. Default: fee_tvl_ratio_24h.
	FeeTVLRatioTW *string `json:"fee_tvl_ratio_tw,omitempty"`
}

// GetGroupParams are optional query parameters for the GetGroup method.
type GetGroupParams struct {
	// Page is the page number (1-based).
	Page *int `json:"page,omitempty"`

	// PageSize is the number of pools to return per page. Max 100.
	PageSize *int `json:"page_size,omitempty"`

	// Query is the search query used to match pools by name, tokens, or address.
	Query *string `json:"query,omitempty"`

	// FilterBy specifies the conditions to filter documents by field values.
	//
	// Format: <expr> [&& <expr> ...]
	// Where each expression is: <field><op><value>
	//
	// Allowed fields:
	// - Numeric: tvl, volume_*, fee_*, fee_tvl_ratio_*, apr_*
	// - Boolean: is_blacklisted
	// - Text: pool_address, name, token_x, token_y
	//
	// Operators:
	// - Numeric: =, >, >=, <, <=
	// - Boolean: =true, =false
	// - Text:
	//   - exact match: =<value>
	//   - multi-value OR: =[value1|value2|...]
	//
	// Examples:
	// - tvl>1000
	// - is_blacklisted=false && volume_24h>=50000
	FilterBy *string `json:"filter_by,omitempty"`

	// SortBy is used to sort results by one or more fields.
	//
	// Format:
	// - Time-windowed metrics: <metric>_<window>:<direction>
	// - Non-windowed metrics: <field>:<direction>
	//
	// - direction: asc or desc
	// - window (when applicable): 5m, 30m, 1h, 2h, 4h, 12h, 24h
	//
	// Available fields:
	// - Time-windowed metrics: volume_*, fee_*, fee_tvl_ratio_*, apr_*
	// - Non-windowed metrics: tvl, fee_pct, bin_step, pool_created_at, farm_apy
	//
	// Default: volume_24h:desc
	//
	// Examples:
	// - volume_24h:desc
	// - fee_1h:asc
	// - tvl:desc
	SortBy *string `json:"sort_by,omitempty"`
}

// TimeframeBasedParams are optional query parameters for time-windowed metrics.
type TimeframeBasedParams struct {
	// Timeframe is the time period interval used for time-windowed metrics.
	// Allowed values: 5m, 30m, 1h, 2h, 4h, 12h, 24h. Default: 24h.
	Timeframe *string `json:"timeframe,omitempty"`

	// StartTime is the Unix timestamp in seconds (inclusive).
	// If omitted, the API uses a default range based on timeframe.
	StartTime *int64 `json:"start_time,omitempty"`

	// EndTime is the Unix timestamp in seconds (inclusive).
	// If omitted, the API uses "now" as the end.
	EndTime *int64 `json:"end_time,omitempty"`
}

// GetPortfolioParams are optional query parameters for the GetPortfolio method.
type GetPortfolioParams struct {
	// User's wallet address.
	User string `json:"user"`

	// Page number for pagination (minimum: 1, default: 1).
	Page *int `json:"page,omitempty"`

	// Page size for pagination (default: 20, maximum: 50).
	PageSize *int `json:"page_size,omitempty"`

	// Only include pools with positions closed within this many days.
	// Applied only for closed positions (minimum: 1, maximum: 365, default: 120).
	DaysBack *int `json:"days_back,omitempty"`
}

// GetOpenPortfolioSort defines the fields to sort open portfolio results by.
type GetOpenPortfolioSort string

const (
	SortByCurrentBalances GetOpenPortfolioSort = "current_balances"
	SortByUnclaimedFee    GetOpenPortfolioSort = "unclaimed_fee"
	SortByFeePerTVL24h    GetOpenPortfolioSort = "fee_per_tvl24h"
)

// SortDirection defines the sort direction.
type SortDirection string

const (
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
)

// GetClosedPositionsParams are optional query parameters for the GetClosedPositions method.
type GetClosedPositionsParams struct {
	// StartTime is the Unix timestamp in seconds (inclusive).
	// If omitted, the API uses a default range based on end time.
	StartTime *int64 `json:"start_time,omitempty"`

	// EndTime is the Unix timestamp in seconds (inclusive).
	// If omitted, the API uses "now" as the end.
	EndTime *int64 `json:"end_time,omitempty"`

	// Limit is the number of positions to return. Default 10, max 100.
	Limit *int `json:"limit,omitempty"`

	// NextCursor is the cursor for pagination.
	NextCursor *string `json:"next_cursor,omitempty"`

	// Pool filters positions to those belonging to a specific pool address.
	Pool *string `json:"pool,omitempty"`
}

// GetOpenPositionsParams are optional query parameters for the GetOpenPositions method.
type GetOpenPositionsParams struct {
	// Pool is an optional comma-separated list of pool addresses to filter positions.
	// Maximum of 50 pool addresses allowed.
	Pool *string `json:"pool,omitempty"`
}

// GetPositionHistoricalEventsParams are optional query parameters for the GetPositionHistoricalEvents method.
type GetPositionHistoricalEventsParams struct {
	// EventType filters by event type (add, remove, claim_fee, claim_reward).
	// If not provided, returns all event types.
	EventType *PositionEventType `json:"event_type,omitempty"`

	// OrderDirection is the sort order for events by block time.
	// asc: Oldest events first, desc: Most recent events first (default).
	OrderDirection *PositionEventOrderDirection `json:"order_direction,omitempty"`
}

// GetPoolPositionPnLParams are query parameters for the GetPoolPositionPnL method.
type GetPoolPositionPnLParams struct {
	// User is the wallet address of the user.
	User string `json:"user"`

	// Status filters positions by status: open, closed, or all (default: all).
	Status *PositionStatus `json:"status,omitempty"`

	// Page is the page number for pagination (minimum: 1, default: 1).
	Page *int `json:"page,omitempty"`

	// PageSize is the page size for pagination (minimum: 1, maximum: 100, default: 20).
	PageSize *int `json:"page_size,omitempty"`
}

// GetOpenPortfolioParams are optional query parameters for the GetOpenPortfolio method.
type GetOpenPortfolioParams struct {
	// User's wallet address.
	User string `json:"user"`

	// Page number for pagination (minimum: 1, default: 1).
	Page *int `json:"page,omitempty"`

	// Page size for pagination (default: 20, maximum: 50).
	PageSize *int `json:"page_size,omitempty"`

	// SortDirection is the sort direction, default is DESC.
	SortDirection *SortDirection `json:"sort_direction,omitempty"`

	// SortBy is the field to sort by, default is current_balances.
	SortBy *GetOpenPortfolioSort `json:"sort_by,omitempty"`
}
