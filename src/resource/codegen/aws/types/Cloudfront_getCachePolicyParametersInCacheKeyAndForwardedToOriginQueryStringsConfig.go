package types

type Cloudfront_getCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig struct {
	// Determines whether any URL query strings in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values are `none`, `whitelist`, `allExcept`, `all`.
	QueryStringBehavior string `json:"queryStringBehavior,omitempty" yaml:"queryStringBehavior,omitempty"`

	// Object that contains a list of query string names. See Items for more information.
	QueryStrings []Cloudfront_getCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryString `json:"queryStrings,omitempty" yaml:"queryStrings,omitempty"`
}
