package types

type Cloudfront_getOriginRequestPolicyHeadersConfig struct {
	//
	HeaderBehavior string `json:"headerBehavior,omitempty" yaml:"headerBehavior,omitempty"`

	//
	Headers []Cloudfront_getOriginRequestPolicyHeadersConfigHeader `json:"headers,omitempty" yaml:"headers,omitempty"`
}
