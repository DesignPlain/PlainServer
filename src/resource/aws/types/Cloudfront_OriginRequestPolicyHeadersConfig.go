package types

type Cloudfront_OriginRequestPolicyHeadersConfig struct {
	//
	HeaderBehavior string `json:"headerBehavior,omitempty" yaml:"headerBehavior,omitempty"`

	//
	Headers Cloudfront_OriginRequestPolicyHeadersConfigHeaders `json:"headers,omitempty" yaml:"headers,omitempty"`
}
