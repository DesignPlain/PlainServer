package types

type Networkfirewall_getFirewallFirewallStatus struct {
	// Aggregated count of all resources used by reference sets in a firewall.
	CapacityUsageSummaries []Networkfirewall_getFirewallFirewallStatusCapacityUsageSummary `json:"capacityUsageSummaries,omitempty" yaml:"capacityUsageSummaries,omitempty"`

	// Summary of sync states for all availability zones in which the firewall is configured.
	ConfigurationSyncStateSummary string `json:"configurationSyncStateSummary,omitempty" yaml:"configurationSyncStateSummary,omitempty"`

	//
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Set of subnets configured for use by the firewall.
	SyncStates []Networkfirewall_getFirewallFirewallStatusSyncState `json:"syncStates,omitempty" yaml:"syncStates,omitempty"`
}
