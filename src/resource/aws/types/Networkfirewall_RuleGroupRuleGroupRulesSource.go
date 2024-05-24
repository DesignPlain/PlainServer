package types

type Networkfirewall_RuleGroupRuleGroupRulesSource struct {
	// A configuration block containing --stateful-- inspection criteria for a domain list rule group. See Rules Source List below for details.
	RulesSourceList Networkfirewall_RuleGroupRuleGroupRulesSourceRulesSourceList `json:"rulesSourceList,omitempty" yaml:"rulesSourceList,omitempty"`

	// The fully qualified name of a file in an S3 bucket that contains Suricata compatible intrusion preventions system (IPS) rules or the Suricata rules as a string. These rules contain --stateful-- inspection criteria and the action to take for traffic that matches the criteria.
	RulesString string `json:"rulesString,omitempty" yaml:"rulesString,omitempty"`

	// Set of configuration blocks containing --stateful-- inspection criteria for 5-tuple rules to be used together in a rule group. See Stateful Rule below for details.
	StatefulRules []Networkfirewall_RuleGroupRuleGroupRulesSourceStatefulRule `json:"statefulRules,omitempty" yaml:"statefulRules,omitempty"`

	// A configuration block containing --stateless-- inspection criteria for a stateless rule group. See Stateless Rules and Custom Actions below for details.
	StatelessRulesAndCustomActions Networkfirewall_RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActions `json:"statelessRulesAndCustomActions,omitempty" yaml:"statelessRulesAndCustomActions,omitempty"`
}
