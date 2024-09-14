package types

type Cloudfront_ResponseHeadersPolicyCustomHeadersConfigItem struct {
	// The value for the HTTP response header.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	//
	Header string `json:"header,omitempty" yaml:"header,omitempty"`

	//
	Override bool `json:"override,omitempty" yaml:"override,omitempty"`
}
