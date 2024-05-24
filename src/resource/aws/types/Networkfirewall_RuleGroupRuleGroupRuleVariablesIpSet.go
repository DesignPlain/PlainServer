package types

type Networkfirewall_RuleGroupRuleGroupRuleVariablesIpSet struct {
	// A configuration block that defines a set of IP addresses. See IP Set below for details.
	IpSet Networkfirewall_RuleGroupRuleGroupRuleVariablesIpSetIpSet `json:"ipSet,omitempty" yaml:"ipSet,omitempty"`

	// A unique alphanumeric string to identify the `ip_set`.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
