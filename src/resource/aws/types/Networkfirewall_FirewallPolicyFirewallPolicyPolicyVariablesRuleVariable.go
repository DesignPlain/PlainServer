package types

type Networkfirewall_FirewallPolicyFirewallPolicyPolicyVariablesRuleVariable struct {
	// A configuration block that defines a set of IP addresses. See IP Set below for details.
	IpSet Networkfirewall_FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableIpSet `json:"ipSet,omitempty" yaml:"ipSet,omitempty"`

	// An alphanumeric string to identify the `ip_set`. Valid values: `HOME_NET`
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
