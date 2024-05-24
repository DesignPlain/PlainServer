package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatement struct {
	// Name of the managed rule group vendor.
	VendorName string `json:"vendorName,omitempty" yaml:"vendorName,omitempty"`

	// Version of the managed rule group. You can set `Version_1.0` or `Version_1.1` etc. If you want to use the default version, do not set anything.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Additional information that's used by a managed rule group. Only one rule attribute is allowed in each config. See `managed_rule_group_configs` for more details
	ManagedRuleGroupConfigs []Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfig `json:"managedRuleGroupConfigs,omitempty" yaml:"managedRuleGroupConfigs,omitempty"`

	// Name of the managed rule group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Action settings to use in the place of the rule actions that are configured inside the rule group. You specify one override for each rule whose action you want to change. See `rule_action_override` below for details.
	RuleActionOverrides []Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverride `json:"ruleActionOverrides,omitempty" yaml:"ruleActionOverrides,omitempty"`

	// Narrows the scope of the statement to matching web requests. This can be any nestable statement, and you can nest statements at any level below this scope-down statement. See `statement` above for details.
	ScopeDownStatement Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatement `json:"scopeDownStatement,omitempty" yaml:"scopeDownStatement,omitempty"`
}
