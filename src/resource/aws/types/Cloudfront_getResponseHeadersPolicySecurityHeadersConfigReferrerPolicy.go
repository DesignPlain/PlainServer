package types

type Cloudfront_getResponseHeadersPolicySecurityHeadersConfigReferrerPolicy struct {
	// Whether CloudFront overrides the X-XSS-Protection HTTP response header received from the origin with the one specified in this response headers policy.
	Override bool `json:"override,omitempty" yaml:"override,omitempty"`

	// Value of the Referrer-Policy HTTP response header. Valid Values: `no-referrer` | `no-referrer-when-downgrade` | `origin` | `origin-when-cross-origin` | `same-origin` | `strict-origin` | `strict-origin-when-cross-origin` | `unsafe-url`
	ReferrerPolicy string `json:"referrerPolicy,omitempty" yaml:"referrerPolicy,omitempty"`
}
