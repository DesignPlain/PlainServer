package ec2clientvpn

type Route struct {
	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId string `json:"clientVpnEndpointId,omitempty" yaml:"clientVpnEndpointId,omitempty"`

	// A brief description of the route.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The IPv4 address range, in CIDR notation, of the route destination.
	DestinationCidrBlock string `json:"destinationCidrBlock,omitempty" yaml:"destinationCidrBlock,omitempty"`

	// The ID of the Subnet to route the traffic through. It must already be attached to the Client VPN.
	TargetVpcSubnetId string `json:"targetVpcSubnetId,omitempty" yaml:"targetVpcSubnetId,omitempty"`
}
