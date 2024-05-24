package types

type Networkfirewall_RuleGroupRuleGroupStatefulRuleOptions struct {
	// Indicates how to manage the order of the rule evaluation for the rule group. Default value: `DEFAULT_ACTION_ORDER`. Valid values: `DEFAULT_ACTION_ORDER`, `STRICT_ORDER`.
	RuleOrder string `json:"ruleOrder,omitempty" yaml:"ruleOrder,omitempty"`
}
