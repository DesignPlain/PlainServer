package types

type Wafv2_RuleGroupRuleCaptchaConfig struct {
	// Defines custom immunity time. See Immunity Time Property below for details.
	ImmunityTimeProperty Wafv2_RuleGroupRuleCaptchaConfigImmunityTimeProperty `json:"immunityTimeProperty,omitempty" yaml:"immunityTimeProperty,omitempty"`
}
