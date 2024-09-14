package types

type Networkfirewall_RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRule struct {
	// A setting that indicates the order in which to run this rule relative to all of the rules that are defined for a stateless rule group. AWS Network Firewall evaluates the rules in a rule group starting with the lowest priority setting.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// A configuration block defining the stateless 5-tuple packet inspection criteria and the action to take on a packet that matches the criteria. See Rule Definition below for details.
	RuleDefinition Networkfirewall_RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRuleRuleDefinition `json:"ruleDefinition,omitempty" yaml:"ruleDefinition,omitempty"`
}
