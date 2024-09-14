package types

type Networkfirewall_getFirewallFirewallStatusCapacityUsageSummary struct {
	// Capacity usage of CIDR blocks used by IP set references in a firewall.
	Cidrs []Networkfirewall_getFirewallFirewallStatusCapacityUsageSummaryCidr `json:"cidrs,omitempty" yaml:"cidrs,omitempty"`
}
