package types

type Networkfirewall_getFirewallFirewallStatusSyncStateAttachment struct {
	// The identifier of the firewall endpoint that AWS Network Firewall has instantiated in the subnet. You use this to identify the firewall endpoint in the VPC route tables, when you redirect the VPC traffic through the endpoint.
	EndpointId string `json:"endpointId,omitempty" yaml:"endpointId,omitempty"`

	//
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// The unique identifier for the subnet.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
