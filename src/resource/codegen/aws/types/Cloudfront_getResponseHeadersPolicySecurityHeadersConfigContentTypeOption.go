package types

type Cloudfront_getResponseHeadersPolicySecurityHeadersConfigContentTypeOption struct {
	// Whether CloudFront overrides the X-XSS-Protection HTTP response header received from the origin with the one specified in this response headers policy.
	Override bool `json:"override,omitempty" yaml:"override,omitempty"`
}
