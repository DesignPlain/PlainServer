package types

type Wafv2_RuleGroupRule struct {
	// Labels to apply to web requests that match the rule match statement. See Rule Label below for details.
	RuleLabels []Wafv2_RuleGroupRuleRuleLabel `json:"ruleLabels,omitempty" yaml:"ruleLabels,omitempty"`

	// The AWS WAF processing statement for the rule, for example `byte_match_statement` or `geo_match_statement`. See Statement below for details.
	Statement Wafv2_RuleGroupRuleStatement `json:"statement,omitempty" yaml:"statement,omitempty"`

	// Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
	VisibilityConfig Wafv2_RuleGroupRuleVisibilityConfig `json:"visibilityConfig,omitempty" yaml:"visibilityConfig,omitempty"`

	// The action that AWS WAF should take on a web request when it matches the rule's statement. Settings at the `aws.wafv2.WebAcl` level can override the rule action setting. See Action below for details.
	Action Wafv2_RuleGroupRuleAction `json:"action,omitempty" yaml:"action,omitempty"`

	// Specifies how AWS WAF should handle CAPTCHA evaluations. See Captcha Configuration below for details.
	CaptchaConfig Wafv2_RuleGroupRuleCaptchaConfig `json:"captchaConfig,omitempty" yaml:"captchaConfig,omitempty"`

	// A friendly name of the rule.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// If you define more than one Rule in a WebACL, AWS WAF evaluates each request against the `rules` in order based on the value of `priority`. AWS WAF processes rules with lower priority first.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`
}
