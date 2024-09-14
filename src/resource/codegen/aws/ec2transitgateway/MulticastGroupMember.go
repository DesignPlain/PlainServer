package ec2transitgateway

type MulticastGroupMember struct {
	// The IP address assigned to the transit gateway multicast group.
	GroupIpAddress string `json:"groupIpAddress,omitempty" yaml:"groupIpAddress,omitempty"`

	// The group members' network interface ID to register with the transit gateway multicast group.
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`

	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId string `json:"transitGatewayMulticastDomainId,omitempty" yaml:"transitGatewayMulticastDomainId,omitempty"`
}
