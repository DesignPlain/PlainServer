package types

type Cloudfront_getOriginRequestPolicyCookiesConfig struct {
	//
	CookieBehavior string `json:"cookieBehavior,omitempty" yaml:"cookieBehavior,omitempty"`

	//
	Cookies []Cloudfront_getOriginRequestPolicyCookiesConfigCookie `json:"cookies,omitempty" yaml:"cookies,omitempty"`
}
