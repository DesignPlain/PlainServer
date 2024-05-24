package types

type Wafv2_WebAclRuleStatementRuleGroupReferenceStatement struct {
	// The Amazon Resource Name (ARN) of the `aws.wafv2.RuleGroup` resource.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// Action settings to use in the place of the rule actions that are configured inside the rule group. You specify one override for each rule whose action you want to change. See `rule_action_override` below for details.
	RuleActionOverrides []Wafv2_WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverride `json:"ruleActionOverrides,omitempty" yaml:"ruleActionOverrides,omitempty"`
}
