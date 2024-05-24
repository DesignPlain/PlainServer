package types

type Cloudfront_getOriginRequestPolicyQueryStringsConfig struct {
	//
	QueryStringBehavior string `json:"queryStringBehavior,omitempty" yaml:"queryStringBehavior,omitempty"`

	//
	QueryStrings []Cloudfront_getOriginRequestPolicyQueryStringsConfigQueryString `json:"queryStrings,omitempty" yaml:"queryStrings,omitempty"`
}
