package types

type Wafv2_WebAclCaptchaConfig struct {
	// Defines custom immunity time. See `immunity_time_property` below for details.
	ImmunityTimeProperty Wafv2_WebAclCaptchaConfigImmunityTimeProperty `json:"immunityTimeProperty,omitempty" yaml:"immunityTimeProperty,omitempty"`
}
