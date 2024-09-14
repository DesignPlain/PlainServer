package types

type Cloudfront_OriginRequestPolicyCookiesConfig struct {
	//
	CookieBehavior string `json:"cookieBehavior,omitempty" yaml:"cookieBehavior,omitempty"`

	//
	Cookies Cloudfront_OriginRequestPolicyCookiesConfigCookies `json:"cookies,omitempty" yaml:"cookies,omitempty"`
}
