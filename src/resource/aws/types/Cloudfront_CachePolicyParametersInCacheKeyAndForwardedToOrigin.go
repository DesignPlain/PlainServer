package types

type Cloudfront_CachePolicyParametersInCacheKeyAndForwardedToOrigin struct {
	// Whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
	CookiesConfig Cloudfront_CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig `json:"cookiesConfig,omitempty" yaml:"cookiesConfig,omitempty"`

	// Flag determines whether the Accept-Encoding HTTP header is included in the cache key and in requests that CloudFront sends to the origin.
	EnableAcceptEncodingBrotli bool `json:"enableAcceptEncodingBrotli,omitempty" yaml:"enableAcceptEncodingBrotli,omitempty"`

	// Whether the Accept-Encoding HTTP header is included in the cache key and in requests sent to the origin by CloudFront.
	EnableAcceptEncodingGzip bool `json:"enableAcceptEncodingGzip,omitempty" yaml:"enableAcceptEncodingGzip,omitempty"`

	// Whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
	HeadersConfig Cloudfront_CachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig `json:"headersConfig,omitempty" yaml:"headersConfig,omitempty"`

	// Whether any URL query strings in viewer requests are included in the cache key. It also automatically includes these query strings in requests that CloudFront sends to the origin. Please refer to the Query String Config for more information.
	QueryStringsConfig Cloudfront_CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig `json:"queryStringsConfig,omitempty" yaml:"queryStringsConfig,omitempty"`
}
