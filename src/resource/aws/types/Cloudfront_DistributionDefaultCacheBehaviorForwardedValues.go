package types

type Cloudfront_DistributionDefaultCacheBehaviorForwardedValues struct {
	// The forwarded values cookies that specifies how CloudFront handles cookies (maximum one).
	Cookies Cloudfront_DistributionDefaultCacheBehaviorForwardedValuesCookies `json:"cookies,omitempty" yaml:"cookies,omitempty"`

	// Headers, if any, that you want CloudFront to vary upon for this cache behavior. Specify `-` to include all headers.
	Headers []string `json:"headers,omitempty" yaml:"headers,omitempty"`

	// Indicates whether you want CloudFront to forward query strings to the origin that is associated with this cache behavior.
	QueryString bool `json:"queryString,omitempty" yaml:"queryString,omitempty"`

	// When specified, along with a value of `true` for `query_string`, all query strings are forwarded, however only the query string keys listed in this argument are cached. When omitted with a value of `true` for `query_string`, all query string keys are cached.
	QueryStringCacheKeys []string `json:"queryStringCacheKeys,omitempty" yaml:"queryStringCacheKeys,omitempty"`
}
