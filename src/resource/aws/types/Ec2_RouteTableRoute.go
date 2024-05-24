package types

type Ec2_RouteTableRoute struct {
	// The Ipv6 CIDR block of the route.
	Ipv6CidrBlock string `json:"ipv6CidrBlock,omitempty" yaml:"ipv6CidrBlock,omitempty"`

	// Identifier of an EC2 network interface.
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`

	// Identifier of an EC2 Transit Gateway.
	TransitGatewayId string `json:"transitGatewayId,omitempty" yaml:"transitGatewayId,omitempty"`

	/*
	   Identifier of a VPC peering connection.

	   Note that the default route, mapping the VPC's CIDR block to "local", is created implicitly and cannot be specified.
	*/
	VpcPeeringConnectionId string `json:"vpcPeeringConnectionId,omitempty" yaml:"vpcPeeringConnectionId,omitempty"`

	// The Amazon Resource Name (ARN) of a core network.
	CoreNetworkArn string `json:"coreNetworkArn,omitempty" yaml:"coreNetworkArn,omitempty"`

	// The CIDR block of the route.
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`

	/*
	   The ID of a managed prefix list destination of the route.

	   One of the following target arguments must be supplied:
	*/
	DestinationPrefixListId string `json:"destinationPrefixListId,omitempty" yaml:"destinationPrefixListId,omitempty"`

	// Identifier of a VPC Egress Only Internet Gateway.
	EgressOnlyGatewayId string `json:"egressOnlyGatewayId,omitempty" yaml:"egressOnlyGatewayId,omitempty"`

	// Identifier of a VPC internet gateway, virtual private gateway, or `local`. `local` routes cannot be created but can be adopted or imported. See the example above.
	GatewayId string `json:"gatewayId,omitempty" yaml:"gatewayId,omitempty"`

	// Identifier of a Outpost local gateway.
	LocalGatewayId string `json:"localGatewayId,omitempty" yaml:"localGatewayId,omitempty"`

	// Identifier of a VPC NAT gateway.
	NatGatewayId string `json:"natGatewayId,omitempty" yaml:"natGatewayId,omitempty"`

	// Identifier of a VPC Endpoint.
	VpcEndpointId string `json:"vpcEndpointId,omitempty" yaml:"vpcEndpointId,omitempty"`

	// Identifier of a carrier gateway. This attribute can only be used when the VPC contains a subnet which is associated with a Wavelength Zone.
	CarrierGatewayId string `json:"carrierGatewayId,omitempty" yaml:"carrierGatewayId,omitempty"`
}
