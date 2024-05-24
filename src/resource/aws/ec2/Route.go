package ec2

type Route struct {
	// The Amazon Resource Name (ARN) of a core network.
	CoreNetworkArn string `json:"coreNetworkArn,omitempty" yaml:"coreNetworkArn,omitempty"`

	// The destination CIDR block.
	DestinationCidrBlock string `json:"destinationCidrBlock,omitempty" yaml:"destinationCidrBlock,omitempty"`

	// Identifier of a VPC Egress Only Internet Gateway.
	EgressOnlyGatewayId string `json:"egressOnlyGatewayId,omitempty" yaml:"egressOnlyGatewayId,omitempty"`

	// Identifier of an EC2 network interface.
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`

	// Identifier of a VPC Endpoint.
	VpcEndpointId string `json:"vpcEndpointId,omitempty" yaml:"vpcEndpointId,omitempty"`

	// Identifier of a carrier gateway. This attribute can only be used when the VPC contains a subnet which is associated with a Wavelength Zone.
	CarrierGatewayId string `json:"carrierGatewayId,omitempty" yaml:"carrierGatewayId,omitempty"`

	/*
	   The ID of a managed prefix list destination.

	   One of the following target arguments must be supplied:
	*/
	DestinationPrefixListId string `json:"destinationPrefixListId,omitempty" yaml:"destinationPrefixListId,omitempty"`

	// Identifier of a Outpost local gateway.
	LocalGatewayId string `json:"localGatewayId,omitempty" yaml:"localGatewayId,omitempty"`

	// The destination IPv6 CIDR block.
	DestinationIpv6CidrBlock string `json:"destinationIpv6CidrBlock,omitempty" yaml:"destinationIpv6CidrBlock,omitempty"`

	// Identifier of a VPC internet gateway or a virtual private gateway. Specify `local` when updating a previously imported local route.
	GatewayId string `json:"gatewayId,omitempty" yaml:"gatewayId,omitempty"`

	// Identifier of a VPC NAT gateway.
	NatGatewayId string `json:"natGatewayId,omitempty" yaml:"natGatewayId,omitempty"`

	/*
	   The ID of the routing table.

	   One of the following destination arguments must be supplied:
	*/
	RouteTableId string `json:"routeTableId,omitempty" yaml:"routeTableId,omitempty"`

	// Identifier of an EC2 Transit Gateway.
	TransitGatewayId string `json:"transitGatewayId,omitempty" yaml:"transitGatewayId,omitempty"`

	/*
	   Identifier of a VPC peering connection.

	   Note that the default route, mapping the VPC's CIDR block to "local", is created implicitly and cannot be specified.
	*/
	VpcPeeringConnectionId string `json:"vpcPeeringConnectionId,omitempty" yaml:"vpcPeeringConnectionId,omitempty"`
}
