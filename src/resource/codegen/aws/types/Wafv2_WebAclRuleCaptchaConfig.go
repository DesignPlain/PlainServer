package types

type Wafv2_WebAclRuleCaptchaConfig struct {
	// Defines custom immunity time. See `immunity_time_property` below for details.
	ImmunityTimeProperty Wafv2_WebAclRuleCaptchaConfigImmunityTimeProperty `json:"immunityTimeProperty,omitempty" yaml:"immunityTimeProperty,omitempty"`
}
