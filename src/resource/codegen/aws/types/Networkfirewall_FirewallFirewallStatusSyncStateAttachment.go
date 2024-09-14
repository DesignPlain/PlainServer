package types

type Networkfirewall_FirewallFirewallStatusSyncStateAttachment struct {
	// The unique identifier of the subnet that you've specified to be used for a firewall endpoint.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// The identifier of the firewall endpoint that AWS Network Firewall has instantiated in the subnet. You use this to identify the firewall endpoint in the VPC route tables, when you redirect the VPC traffic through the endpoint.
	EndpointId string `json:"endpointId,omitempty" yaml:"endpointId,omitempty"`
}
