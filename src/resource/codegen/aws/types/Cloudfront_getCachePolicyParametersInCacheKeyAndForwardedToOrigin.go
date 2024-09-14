package types

type Cloudfront_getCachePolicyParametersInCacheKeyAndForwardedToOrigin struct {
	// Object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
	CookiesConfigs []Cloudfront_getCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig `json:"cookiesConfigs,omitempty" yaml:"cookiesConfigs,omitempty"`

	// A flag that can affect whether the Accept-Encoding HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.
	EnableAcceptEncodingBrotli bool `json:"enableAcceptEncodingBrotli,omitempty" yaml:"enableAcceptEncodingBrotli,omitempty"`

	// A flag that can affect whether the Accept-Encoding HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.
	EnableAcceptEncodingGzip bool `json:"enableAcceptEncodingGzip,omitempty" yaml:"enableAcceptEncodingGzip,omitempty"`

	// Object that determines whether any HTTP headers (and if so, which headers) are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
	HeadersConfigs []Cloudfront_getCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig `json:"headersConfigs,omitempty" yaml:"headersConfigs,omitempty"`

	// Object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Query String Config for more information.
	QueryStringsConfigs []Cloudfront_getCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig `json:"queryStringsConfigs,omitempty" yaml:"queryStringsConfigs,omitempty"`
}
