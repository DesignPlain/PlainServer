package types

type Compute_BackendBucketCdnPolicyCacheKeyPolicy struct {
	/*
	   Names of query string parameters to include in cache keys.
	   Default parameters are always included. '&' and '=' will
	   be percent encoded and not treated as delimiters.
	*/
	QueryStringWhitelists []string `json:"queryStringWhitelists,omitempty" yaml:"queryStringWhitelists,omitempty"`

	/*
	   Allows HTTP request headers (by name) to be used in the
	   cache key.
	*/
	IncludeHttpHeaders []string `json:"includeHttpHeaders,omitempty" yaml:"includeHttpHeaders,omitempty"`
}
