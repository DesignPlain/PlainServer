package types

type Networkfirewall_RuleGroupRuleGroupRuleVariables struct {
	// Set of configuration blocks that define IP address information. See IP Sets below for details.
	IpSets []Networkfirewall_RuleGroupRuleGroupRuleVariablesIpSet `json:"ipSets,omitempty" yaml:"ipSets,omitempty"`

	// Set of configuration blocks that define port range information. See Port Sets below for details.
	PortSets []Networkfirewall_RuleGroupRuleGroupRuleVariablesPortSet `json:"portSets,omitempty" yaml:"portSets,omitempty"`
}
