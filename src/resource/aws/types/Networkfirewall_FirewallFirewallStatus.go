package types

type Networkfirewall_FirewallFirewallStatus struct {
	// Set of subnets configured for use by the firewall.
	SyncStates []Networkfirewall_FirewallFirewallStatusSyncState `json:"syncStates,omitempty" yaml:"syncStates,omitempty"`
}
