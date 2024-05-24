package types

type Networkfirewall_RuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActionsStatelessRuleRuleDefinitionMatchAttributesSourcePort struct {
	// The lower limit of the port range. This must be less than or equal to the `to_port`.
	FromPort int `json:"fromPort,omitempty" yaml:"fromPort,omitempty"`

	// The upper limit of the port range. This must be greater than or equal to the `from_port`.
	ToPort int `json:"toPort,omitempty" yaml:"toPort,omitempty"`
}
