package types

type Cloudfront_ResponseHeadersPolicyServerTimingHeadersConfig struct {
	// A Whether CloudFront adds the `Server-Timing` header to HTTP responses that it sends in response to requests that match a cache behavior that's associated with this response headers policy.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// A number 0–100 (inclusive) that specifies the percentage of responses that you want CloudFront to add the Server-Timing header to. Valid range: Minimum value of 0.0. Maximum value of 100.0.
	SamplingRate float64 `json:"samplingRate,omitempty" yaml:"samplingRate,omitempty"`
}
