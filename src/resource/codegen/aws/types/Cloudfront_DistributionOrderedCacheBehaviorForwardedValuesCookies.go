package types

type Cloudfront_DistributionOrderedCacheBehaviorForwardedValuesCookies struct {
	// Whether you want CloudFront to forward cookies to the origin that is associated with this cache behavior. You can specify `all`, `none` or `whitelist`. If `whitelist`, you must include the subsequent `whitelisted_names`.
	Forward string `json:"forward,omitempty" yaml:"forward,omitempty"`

	// If you have specified `whitelist` to `forward`, the whitelisted cookies that you want CloudFront to forward to your origin.
	WhitelistedNames []string `json:"whitelistedNames,omitempty" yaml:"whitelistedNames,omitempty"`
}
