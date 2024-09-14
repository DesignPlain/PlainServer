package types

type Wafregional_WebAclRule struct {
	// Configuration block of the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Not used if `type` is `GROUP`. Detailed below.
	Action Wafregional_WebAclRuleAction `json:"action,omitempty" yaml:"action,omitempty"`

	// Configuration block of the override the action that a group requests CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Only used if `type` is `GROUP`. Detailed below.
	OverrideAction Wafregional_WebAclRuleOverrideAction `json:"overrideAction,omitempty" yaml:"overrideAction,omitempty"`

	/*
	   Specifies the order in which the rules in a WebACL are evaluated.
	   Rules with a lower value are evaluated before rules with a higher value.
	*/
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// ID of the associated WAF (Regional) rule (e.g., `aws.wafregional.Rule`). WAF (Global) rules cannot be used.
	RuleId string `json:"ruleId,omitempty" yaml:"ruleId,omitempty"`

	// The rule type, either `REGULAR`, as defined by [Rule](http://docs.aws.amazon.com/waf/latest/APIReference/API_Rule.html), `RATE_BASED`, as defined by [RateBasedRule](http://docs.aws.amazon.com/waf/latest/APIReference/API_RateBasedRule.html), or `GROUP`, as defined by [RuleGroup](https://docs.aws.amazon.com/waf/latest/APIReference/API_RuleGroup.html). The default is REGULAR. If you add a RATE_BASED rule, you need to set `type` as `RATE_BASED`. If you add a GROUP rule, you need to set `type` as `GROUP`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
