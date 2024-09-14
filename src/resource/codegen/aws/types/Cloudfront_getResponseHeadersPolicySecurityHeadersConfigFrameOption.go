package types

type Cloudfront_getResponseHeadersPolicySecurityHeadersConfigFrameOption struct {
	// Value of the X-Frame-Options HTTP response header. Valid values: `DENY` | `SAMEORIGIN`
	FrameOption string `json:"frameOption,omitempty" yaml:"frameOption,omitempty"`

	// Whether CloudFront overrides the X-XSS-Protection HTTP response header received from the origin with the one specified in this response headers policy.
	Override bool `json:"override,omitempty" yaml:"override,omitempty"`
}
