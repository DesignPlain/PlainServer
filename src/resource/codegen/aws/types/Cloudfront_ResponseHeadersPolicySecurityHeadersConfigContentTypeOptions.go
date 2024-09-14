package types

type Cloudfront_ResponseHeadersPolicySecurityHeadersConfigContentTypeOptions struct {
	// Whether CloudFront overrides the `X-Content-Type-Options` HTTP response header received from the origin with the one specified in this response headers policy.
	Override bool `json:"override,omitempty" yaml:"override,omitempty"`
}
