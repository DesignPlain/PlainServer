package types

type Cloudfront_getResponseHeadersPolicySecurityHeadersConfigXssProtection struct {
	// Whether CloudFront includes the mode=block directive in the X-XSS-Protection header.
	ModeBlock bool `json:"modeBlock,omitempty" yaml:"modeBlock,omitempty"`

	// Whether CloudFront overrides the X-XSS-Protection HTTP response header received from the origin with the one specified in this response headers policy.
	Override bool `json:"override,omitempty" yaml:"override,omitempty"`

	// Boolean value that determines the value of the X-XSS-Protection HTTP response header. When this setting is true, the value of the X-XSS-Protection header is 1. When this setting is false, the value of the X-XSS-Protection header is 0.
	Protection bool `json:"protection,omitempty" yaml:"protection,omitempty"`

	// Whether CloudFront sets a reporting URI in the X-XSS-Protection header.
	ReportUri string `json:"reportUri,omitempty" yaml:"reportUri,omitempty"`
}
