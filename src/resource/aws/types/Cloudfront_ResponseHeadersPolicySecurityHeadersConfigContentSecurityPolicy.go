package types

type Cloudfront_ResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy struct {
	// The policy directives and their values that CloudFront includes as values for the `Content-Security-Policy` HTTP response header.
	ContentSecurityPolicy string `json:"contentSecurityPolicy,omitempty" yaml:"contentSecurityPolicy,omitempty"`

	// Whether CloudFront overrides the `Content-Security-Policy` HTTP response header received from the origin with the one specified in this response headers policy.
	Override bool `json:"override,omitempty" yaml:"override,omitempty"`
}
