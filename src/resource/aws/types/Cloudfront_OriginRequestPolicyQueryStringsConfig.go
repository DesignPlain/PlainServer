package types

type Cloudfront_OriginRequestPolicyQueryStringsConfig struct {
	//
	QueryStringBehavior string `json:"queryStringBehavior,omitempty" yaml:"queryStringBehavior,omitempty"`

	//
	QueryStrings Cloudfront_OriginRequestPolicyQueryStringsConfigQueryStrings `json:"queryStrings,omitempty" yaml:"queryStrings,omitempty"`
}
