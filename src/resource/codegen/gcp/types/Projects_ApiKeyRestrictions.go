package types

type Projects_ApiKeyRestrictions struct {
	// The Android apps that are allowed to use the key.
	AndroidKeyRestrictions Projects_ApiKeyRestrictionsAndroidKeyRestrictions `json:"androidKeyRestrictions,omitempty" yaml:"androidKeyRestrictions,omitempty"`

	// A restriction for a specific service and optionally one or more specific methods. Requests are allowed if they match any of these restrictions. If no restrictions are specified, all targets are allowed.
	ApiTargets []Projects_ApiKeyRestrictionsApiTarget `json:"apiTargets,omitempty" yaml:"apiTargets,omitempty"`

	// The HTTP referrers (websites) that are allowed to use the key.
	BrowserKeyRestrictions Projects_ApiKeyRestrictionsBrowserKeyRestrictions `json:"browserKeyRestrictions,omitempty" yaml:"browserKeyRestrictions,omitempty"`

	// The iOS apps that are allowed to use the key.
	IosKeyRestrictions Projects_ApiKeyRestrictionsIosKeyRestrictions `json:"iosKeyRestrictions,omitempty" yaml:"iosKeyRestrictions,omitempty"`

	// The IP addresses of callers that are allowed to use the key.
	ServerKeyRestrictions Projects_ApiKeyRestrictionsServerKeyRestrictions `json:"serverKeyRestrictions,omitempty" yaml:"serverKeyRestrictions,omitempty"`
}
