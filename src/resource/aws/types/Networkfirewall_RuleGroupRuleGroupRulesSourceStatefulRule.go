package types

type Networkfirewall_RuleGroupRuleGroupRulesSourceStatefulRule struct {
	// Action to take with packets in a traffic flow when the flow matches the stateful rule criteria. For all actions, AWS Network Firewall performs the specified action and discontinues stateful inspection of the traffic flow. Valid values: `ALERT`, `DROP`, `PASS`, or `REJECT`.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// A configuration block containing the stateful 5-tuple inspection criteria for the rule, used to inspect traffic flows. See Header below for details.
	Header Networkfirewall_RuleGroupRuleGroupRulesSourceStatefulRuleHeader `json:"header,omitempty" yaml:"header,omitempty"`

	// Set of configuration blocks containing additional settings for a stateful rule. See Rule Option below for details.
	RuleOptions []Networkfirewall_RuleGroupRuleGroupRulesSourceStatefulRuleRuleOption `json:"ruleOptions,omitempty" yaml:"ruleOptions,omitempty"`
}
