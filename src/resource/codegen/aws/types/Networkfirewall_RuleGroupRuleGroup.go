package types

type Networkfirewall_RuleGroupRuleGroup struct {
	// A configuration block that defines additional settings available to use in the rules defined in the rule group. Can only be specified for --stateful-- rule groups. See Rule Variables below for details.
	RuleVariables Networkfirewall_RuleGroupRuleGroupRuleVariables `json:"ruleVariables,omitempty" yaml:"ruleVariables,omitempty"`

	// A configuration block that defines the stateful or stateless rules for the rule group. See Rules Source below for details.
	RulesSource Networkfirewall_RuleGroupRuleGroupRulesSource `json:"rulesSource,omitempty" yaml:"rulesSource,omitempty"`

	// A configuration block that defines stateful rule options for the rule group. See Stateful Rule Options below for details.
	StatefulRuleOptions Networkfirewall_RuleGroupRuleGroupStatefulRuleOptions `json:"statefulRuleOptions,omitempty" yaml:"statefulRuleOptions,omitempty"`

	// A configuration block that defines the IP Set References for the rule group. See Reference Sets below for details. Please notes that there can only be a maximum of 5 `reference_sets` in a `rule_group`. See the [AWS documentation](https://docs.aws.amazon.com/network-firewall/latest/developerguide/rule-groups-ip-set-references.html#rule-groups-ip-set-reference-limits) for details.
	ReferenceSets Networkfirewall_RuleGroupRuleGroupReferenceSets `json:"referenceSets,omitempty" yaml:"referenceSets,omitempty"`
}
