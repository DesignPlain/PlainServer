package types

type Compute_getBackendServiceCdnPolicyCacheKeyPolicy struct {
	// If true, http and https requests will be cached separately.
	IncludeProtocol bool `json:"includeProtocol,omitempty" yaml:"includeProtocol,omitempty"`

	/*
	   If true, include query string parameters in the cache key
	   according to query_string_whitelist and
	   query_string_blacklist. If neither is set, the entire query
	   string will be included.

	   If false, the query string will be excluded from the cache
	   key entirely.
	*/
	IncludeQueryString bool `json:"includeQueryString,omitempty" yaml:"includeQueryString,omitempty"`

	/*
	   Names of query string parameters to exclude in cache keys.

	   All other parameters will be included. Either specify
	   query_string_whitelist or query_string_blacklist, not both.
	   '&' and '=' will be percent encoded and not treated as
	   delimiters.
	*/
	QueryStringBlacklists []string `json:"queryStringBlacklists,omitempty" yaml:"queryStringBlacklists,omitempty"`

	/*
	   Names of query string parameters to include in cache keys.

	   All other parameters will be excluded. Either specify
	   query_string_whitelist or query_string_blacklist, not both.
	   '&' and '=' will be percent encoded and not treated as
	   delimiters.
	*/
	QueryStringWhitelists []string `json:"queryStringWhitelists,omitempty" yaml:"queryStringWhitelists,omitempty"`

	// If true requests to different hosts will be cached separately.
	IncludeHost bool `json:"includeHost,omitempty" yaml:"includeHost,omitempty"`

	/*
	   Allows HTTP request headers (by name) to be used in the
	   cache key.
	*/
	IncludeHttpHeaders []string `json:"includeHttpHeaders,omitempty" yaml:"includeHttpHeaders,omitempty"`

	// Names of cookies to include in cache keys.
	IncludeNamedCookies []string `json:"includeNamedCookies,omitempty" yaml:"includeNamedCookies,omitempty"`
}
