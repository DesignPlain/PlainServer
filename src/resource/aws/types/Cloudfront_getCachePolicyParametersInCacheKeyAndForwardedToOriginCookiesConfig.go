package types

type Cloudfront_getCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig struct {
	// Determines whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin. Valid values are `none`, `whitelist`, `allExcept`, `all`.
	CookieBehavior string `json:"cookieBehavior,omitempty" yaml:"cookieBehavior,omitempty"`

	// Object that contains a list of cookie names. See Items for more information.
	Cookies []Cloudfront_getCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookie `json:"cookies,omitempty" yaml:"cookies,omitempty"`
}
