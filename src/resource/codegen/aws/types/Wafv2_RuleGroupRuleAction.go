package types

type Wafv2_RuleGroupRuleAction struct {
	// Instructs AWS WAF to run a `CAPTCHA` check against the web request. See Captcha below for details.
	Captcha Wafv2_RuleGroupRuleActionCaptcha `json:"captcha,omitempty" yaml:"captcha,omitempty"`

	// Instructs AWS WAF to run a check against the request to verify that the request is coming from a legitimate client session. See Challenge below for details.
	Challenge Wafv2_RuleGroupRuleActionChallenge `json:"challenge,omitempty" yaml:"challenge,omitempty"`

	// Instructs AWS WAF to count the web request and allow it. See Count below for details.
	Count Wafv2_RuleGroupRuleActionCount `json:"count,omitempty" yaml:"count,omitempty"`

	// Instructs AWS WAF to allow the web request. See Allow below for details.
	Allow Wafv2_RuleGroupRuleActionAllow `json:"allow,omitempty" yaml:"allow,omitempty"`

	// Instructs AWS WAF to block the web request. See Block below for details.
	Block Wafv2_RuleGroupRuleActionBlock `json:"block,omitempty" yaml:"block,omitempty"`
}
