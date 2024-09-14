package types

type Networkfirewall_FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableIpSet struct {
	// Set of IPv4 or IPv6 addresses in CIDR notation to use for the Suricata `HOME_NET` variable.
	Definitions []string `json:"definitions,omitempty" yaml:"definitions,omitempty"`
}
