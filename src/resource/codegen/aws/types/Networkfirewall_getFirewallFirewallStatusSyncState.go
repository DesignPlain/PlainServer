package types

type Networkfirewall_getFirewallFirewallStatusSyncState struct {
	// Nested list describing the attachment status of the firewall's association with a single VPC subnet.
	Attachments []Networkfirewall_getFirewallFirewallStatusSyncStateAttachment `json:"attachments,omitempty" yaml:"attachments,omitempty"`

	// The Availability Zone where the subnet is configured.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`
}
