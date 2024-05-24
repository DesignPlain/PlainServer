package types

type Cloudfront_CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig struct {
	// Whether URL query strings in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values for `query_string_behavior` are `none`, `whitelist`, `allExcept`, and `all`.
	QueryStringBehavior string `json:"queryStringBehavior,omitempty" yaml:"queryStringBehavior,omitempty"`

	// Configuration parameter that contains a list of query string names. See Items for more information.
	QueryStrings Cloudfront_CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings `json:"queryStrings,omitempty" yaml:"queryStrings,omitempty"`
}
