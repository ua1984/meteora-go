package dammv2

// ListPoolsParams are optional query parameters for the ListPools method.
type ListPoolsParams struct {
	// Page is the page number (1-based).
	Page *int `json:"page,omitempty"`

	// PageSize is the number of pools to return per page. Max 1000.
	PageSize *int `json:"page_size,omitempty"`

	// SortBy sorts results by one or more fields.
	//
	// Format:
	//   - Time-windowed metrics: <metric>_<window>:<direction>
	//   - Non-windowed metrics: <field>:<direction>
	//
	// direction: asc or desc
	// window (when applicable): 5m 30m 1h 2h 4h 12h 24h
	//
	// Available fields:
	//   - Time-windowed: volume_* fee_* fee_tvl_ratio_* apr_*
	//   - Non-windowed: tvl fee_pct bin_step pool_created_at farm_apy
	//
	// Default: volume_24h:desc
	SortBy *string `json:"sort_by,omitempty"`

	// Query is a search query used to match pools by name, tokens, or address.
	Query *string `json:"query,omitempty"`

	// FilterBy specifies conditions to filter pools by field values.
	//
	// Format: <expr> [&& <expr> ...]
	// Where each expression is: <field><op><value>
	//
	// Allowed fields:
	//   - Numeric: tvl volume_* fee_* fee_tvl_ratio_* apr_*
	//   - Boolean: is_blacklisted
	//   - Text: pool_address name token_x token_y
	//
	// Example: "is_blacklisted=false && volume_24h>=50000"
	FilterBy *string `json:"filter_by,omitempty"`
}

// ListGroupsParams are optional query parameters for the ListGroups method.
type ListGroupsParams struct {
	// Page is the page number (1-based).
	Page *int `json:"page,omitempty"`

	// PageSize is the number of groups to return per page. Max 100.
	PageSize *int `json:"page_size,omitempty"`

	// SortBy sorts results by one or more fields.
	//
	// Format:
	//   - Time-windowed metrics: <metric>_<window>:<direction>
	//   - Non-windowed metrics: <field>:<direction>
	//
	// direction: asc or desc
	// window (when applicable): 5m 30m 1h 2h 4h 12h 24h
	//
	// Default: volume_24h:desc
	SortBy *string `json:"sort_by,omitempty"`

	// Query is a search query used to match pools by name, tokens, or address.
	Query *string `json:"query,omitempty"`

	// FilterBy specifies conditions to filter groups by field values.
	//
	// Format: <expr> [&& <expr> ...]
	// Where each expression is: <field><op><value>
	//
	// Example: "is_blacklisted=false && volume_24h>=50000"
	FilterBy *string `json:"filter_by,omitempty"`

	// VolumeTW is the time window to aggregate volume. Returns sum.
	// Default: volume_24h.
	VolumeTW *string `json:"volume_tw,omitempty"`

	// FeeTVLRatioTW is the time window to aggregate fee TVL ratio. Returns max.
	// Default: fee_tvl_ratio_24h.
	FeeTVLRatioTW *string `json:"fee_tvl_ratio_tw,omitempty"`
}

// GetClosedPositionsParams are optional query parameters for the GetClosedPositions method.
type GetClosedPositionsParams struct {
	// StartTime is the Unix timestamp (seconds, inclusive) for the start of the time range.
	StartTime *int64 `json:"start_time,omitempty"`

	// EndTime is the Unix timestamp (seconds, inclusive) for the end of the time range.
	EndTime *int64 `json:"end_time,omitempty"`

	// Limit is the number of positions to return. Default 10, max 100.
	Limit *int `json:"limit,omitempty"`

	// NextCursor is the pagination cursor returned by a previous call.
	NextCursor *string `json:"next_cursor,omitempty"`

	// Pool filters results to positions in the specified pool address.
	Pool *string `json:"pool,omitempty"`
}

// GetOpenPositionsParams are optional query parameters for the GetOpenPositions method.
type GetOpenPositionsParams struct {
	// Pool is an optional comma-separated list of pool addresses to filter positions.
	// Maximum of 50 pool addresses allowed.
	Pool *string `json:"pool,omitempty"`
}

// GetGroupParams are optional query parameters for the GetGroup method.
type GetGroupParams struct {
	// Page is the page number (1-based).
	Page *int `json:"page,omitempty"`

	// PageSize is the number of pools to return per page. Max 100.
	PageSize *int `json:"page_size,omitempty"`

	// SortBy sorts results by one or more fields.
	//
	// Format:
	//   - Time-windowed metrics: <metric>_<window>:<direction>
	//   - Non-windowed metrics: <field>:<direction>
	//
	// direction: asc or desc
	// window (when applicable): 5m 30m 1h 2h 4h 12h 24h
	//
	// Default: volume_24h:desc
	SortBy *string `json:"sort_by,omitempty"`

	// Query is a search query used to match pools by name, tokens, or address.
	Query *string `json:"query,omitempty"`

	// FilterBy specifies conditions to filter pools by field values.
	//
	// Format: <expr> [&& <expr> ...]
	// Where each expression is: <field><op><value>
	//
	// Example: "is_blacklisted=false && volume_24h>=50000"
	FilterBy *string `json:"filter_by,omitempty"`
}
