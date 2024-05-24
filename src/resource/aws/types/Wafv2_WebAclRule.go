package types

type Wafv2_WebAclRule struct {
	// Override action to apply to the rules in a rule group. Used only for rule --statements that reference a rule group--, like `rule_group_reference_statement` and `managed_rule_group_statement`. See `override_action` below for details.
	OverrideAction Wafv2_WebAclRuleOverrideAction `json:"overrideAction,omitempty" yaml:"overrideAction,omitempty"`

	// If you define more than one Rule in a WebACL, AWS WAF evaluates each request against the `rules` in order based on the value of `priority`. AWS WAF processes rules with lower priority first.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// Labels to apply to web requests that match the rule match statement. See `rule_label` below for details.
	RuleLabels []Wafv2_WebAclRuleRuleLabel `json:"ruleLabels,omitempty" yaml:"ruleLabels,omitempty"`

	// The AWS WAF processing statement for the rule, for example `byte_match_statement` or `geo_match_statement`. See `statement` below for details.
	Statement Wafv2_WebAclRuleStatement `json:"statement,omitempty" yaml:"statement,omitempty"`

	// Defines and enables Amazon CloudWatch metrics and web request sample collection. See `visibility_config` below for details.
	VisibilityConfig Wafv2_WebAclRuleVisibilityConfig `json:"visibilityConfig,omitempty" yaml:"visibilityConfig,omitempty"`

	// Action that AWS WAF should take on a web request when it matches the rule's statement. This is used only for rules whose --statements do not reference a rule group--. See `action` for details.
	Action Wafv2_WebAclRuleAction `json:"action,omitempty" yaml:"action,omitempty"`

	// Specifies how AWS WAF should handle CAPTCHA evaluations. See `captcha_config` below for details.
	CaptchaConfig Wafv2_WebAclRuleCaptchaConfig `json:"captchaConfig,omitempty" yaml:"captchaConfig,omitempty"`

	// Friendly name of the rule. Note that the provider assumes that rules with names matching this pattern, `^ShieldMitigationRuleGroup_<account-id>_<web-acl-guid>_.-`, are AWS-added for [automatic application layer DDoS mitigation activities](https://docs.aws.amazon.com/waf/latest/developerguide/ddos-automatic-app-layer-response-rg.html). Such rules will be ignored by the provider unless you explicitly include them in your configuration (for example, by using the AWS CLI to discover their properties and creating matching configuration). However, since these rules are owned and managed by AWS, you may get permission errors.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
