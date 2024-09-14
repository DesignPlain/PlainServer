package types

type Cloudfront_CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig struct {
	// Whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values for `cookie_behavior` are `none`, `whitelist`, `allExcept`, and `all`.
	CookieBehavior string `json:"cookieBehavior,omitempty" yaml:"cookieBehavior,omitempty"`

	// Object that contains a list of cookie names. See Items for more information.
	Cookies Cloudfront_CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies `json:"cookies,omitempty" yaml:"cookies,omitempty"`
}
