package dammv2

// ListPoolsParams are optional query parameters for the ListPools method.
type ListPoolsParams struct {
	// Page is the 1-based page number to retrieve.
	Page *int `json:"page,omitempty"`

	// Limit is the number of pools per page.
	Limit *int `json:"limit,omitempty"`

	// SortBy is the field name to sort results by (e.g., "tvl", "volume", "fees").
	SortBy *string `json:"sort_by,omitempty"`

	// SortOrder is the sort direction ("asc" or "desc").
	SortOrder *string `json:"sort_order,omitempty"`

	// SearchTerm filters pools by name, symbol, or address substring.
	SearchTerm *string `json:"search_term,omitempty"`

	// HideBlacklist excludes blacklisted pools from results when true.
	HideBlacklist *bool `json:"hide_blacklist,omitempty"`

	// IncludeTags filters pools to only include those with the specified tag.
	IncludeTags *string `json:"include_tags,omitempty"`

	// ExcludeTags filters pools to exclude those with the specified tag.
	ExcludeTags *string `json:"exclude_tags,omitempty"`
}

// ListGroupsParams are optional query parameters for the ListGroups method.
type ListGroupsParams struct {
	// Page is the 1-based page number to retrieve.
	Page *int `json:"page,omitempty"`

	// Limit is the number of groups per page.
	Limit *int `json:"limit,omitempty"`

	// SortBy is the field name to sort results by.
	SortBy *string `json:"sort_by,omitempty"`

	// SortOrder is the sort direction ("asc" or "desc").
	SortOrder *string `json:"sort_order,omitempty"`

	// SearchTerm filters groups by name or token address substring.
	SearchTerm *string `json:"search_term,omitempty"`

	// HideBlacklist excludes groups containing only blacklisted pools when true.
	HideBlacklist *bool `json:"hide_blacklist,omitempty"`
}

// GetGroupParams are optional query parameters for the GetGroup method.
type GetGroupParams struct {
	// Page is the 1-based page number to retrieve.
	Page *int `json:"page,omitempty"`

	// Limit is the number of pools per page within the group.
	Limit *int `json:"limit,omitempty"`

	// SortBy is the field name to sort pools by within the group.
	SortBy *string `json:"sort_by,omitempty"`

	// SortOrder is the sort direction ("asc" or "desc").
	SortOrder *string `json:"sort_order,omitempty"`
}
