package types

type Cloudfront_getCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig struct {
	// Determines whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values are `none`, `whitelist`.
	HeaderBehavior string `json:"headerBehavior,omitempty" yaml:"headerBehavior,omitempty"`

	// Object that contains a list of header names. See Items for more information.
	Headers []Cloudfront_getCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeader `json:"headers,omitempty" yaml:"headers,omitempty"`
}
