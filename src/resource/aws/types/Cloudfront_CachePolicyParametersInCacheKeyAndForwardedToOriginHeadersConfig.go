package types

type Cloudfront_CachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig struct {
	// Whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values for `header_behavior` are `none` and `whitelist`.
	HeaderBehavior string `json:"headerBehavior,omitempty" yaml:"headerBehavior,omitempty"`

	// Object contains a list of header names. See Items for more information.
	Headers Cloudfront_CachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders `json:"headers,omitempty" yaml:"headers,omitempty"`
}
