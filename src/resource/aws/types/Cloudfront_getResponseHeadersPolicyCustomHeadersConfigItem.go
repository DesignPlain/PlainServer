package types

type Cloudfront_getResponseHeadersPolicyCustomHeadersConfigItem struct {
	// The HTTP header name.
	Header string `json:"header,omitempty" yaml:"header,omitempty"`

	// Whether CloudFront overrides the X-XSS-Protection HTTP response header received from the origin with the one specified in this response headers policy.
	Override bool `json:"override,omitempty" yaml:"override,omitempty"`

	// Value for the HTTP response header.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
