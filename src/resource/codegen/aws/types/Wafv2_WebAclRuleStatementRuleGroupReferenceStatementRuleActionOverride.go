package types

type Wafv2_WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverride struct {
	// Override action to use, in place of the configured action of the rule in the rule group. See `action` for details.
	ActionToUse Wafv2_WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUse `json:"actionToUse,omitempty" yaml:"actionToUse,omitempty"`

	// Name of the rule to override. See the [documentation](https://docs.aws.amazon.com/waf/latest/developerguide/aws-managed-rule-groups-list.html) for a list of names in the appropriate rule group in use.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
