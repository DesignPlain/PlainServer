package types

type Lightsail_DistributionCacheBehaviorSettingsForwardedCookies struct {
	// The specific cookies to forward to your distribution's origin.
	CookiesAllowLists []string `json:"cookiesAllowLists,omitempty" yaml:"cookiesAllowLists,omitempty"`

	// Specifies which cookies to forward to the distribution's origin for a cache behavior: all, none, or allow-list to forward only the cookies specified in the cookiesAllowList parameter.
	Option string `json:"option,omitempty" yaml:"option,omitempty"`
}
