package types

type Networkfirewall_getFirewallFirewallStatusCapacityUsageSummaryCidr struct {
	// Number of CIDR blocks used by the IP set references in a firewall.
	UtilizedCidrCount int `json:"utilizedCidrCount,omitempty" yaml:"utilizedCidrCount,omitempty"`

	// Available number of CIDR blocks available for use by the IP set references in a firewall.
	AvailableCidrCount int `json:"availableCidrCount,omitempty" yaml:"availableCidrCount,omitempty"`

	// The list of IP set references used by a firewall.
	IpSetReferences []Networkfirewall_getFirewallFirewallStatusCapacityUsageSummaryCidrIpSetReference `json:"ipSetReferences,omitempty" yaml:"ipSetReferences,omitempty"`
}
