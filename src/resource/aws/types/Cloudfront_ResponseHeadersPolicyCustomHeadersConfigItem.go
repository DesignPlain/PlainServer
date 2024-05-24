package types

type Cloudfront_ResponseHeadersPolicyCustomHeadersConfigItem struct {
	// The HTTP response header name.
	Header string `json:"header,omitempty" yaml:"header,omitempty"`

	// Whether CloudFront overrides the `Content-Security-Policy` HTTP response header received from the origin with the one specified in this response headers policy.
	Override bool `json:"override,omitempty" yaml:"override,omitempty"`

	// The value for the HTTP response header.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
