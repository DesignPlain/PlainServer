package types

type Cloudfront_getResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy struct {
	// The policy directives and their values that CloudFront includes as values for the Content-Security-Policy HTTP response header.
	ContentSecurityPolicy string `json:"contentSecurityPolicy,omitempty" yaml:"contentSecurityPolicy,omitempty"`

	// Whether CloudFront overrides the X-XSS-Protection HTTP response header received from the origin with the one specified in this response headers policy.
	Override bool `json:"override,omitempty" yaml:"override,omitempty"`
}
