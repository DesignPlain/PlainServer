package types

type Cloudfront_getResponseHeadersPolicyServerTimingHeadersConfig struct {
	// Whether CloudFront adds the `Server-Timing` header to HTTP responses that it sends in response to requests that match a cache behavior that's associated with this response headers policy.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Number 0–100 (inclusive) that specifies the percentage of responses that you want CloudFront to add the Server-Timing header to.
	SamplingRate float64 `json:"samplingRate,omitempty" yaml:"samplingRate,omitempty"`
}
