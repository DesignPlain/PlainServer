package types

type Networkfirewall_RuleGroupRuleGroupRuleVariablesPortSet struct {
	// An unique alphanumeric string to identify the `port_set`.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// A configuration block that defines a set of port ranges. See Port Set below for details.
	PortSet Networkfirewall_RuleGroupRuleGroupRuleVariablesPortSetPortSet `json:"portSet,omitempty" yaml:"portSet,omitempty"`
}
