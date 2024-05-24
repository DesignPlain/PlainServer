package types

type Ec2_getRouteTableRoute struct {
	// ID of the Egress Only Internet Gateway.
	EgressOnlyGatewayId string `json:"egressOnlyGatewayId,omitempty" yaml:"egressOnlyGatewayId,omitempty"`

	// Local Gateway ID.
	LocalGatewayId string `json:"localGatewayId,omitempty" yaml:"localGatewayId,omitempty"`

	// NAT Gateway ID.
	NatGatewayId string `json:"natGatewayId,omitempty" yaml:"natGatewayId,omitempty"`

	// ARN of the core network.
	CoreNetworkArn string `json:"coreNetworkArn,omitempty" yaml:"coreNetworkArn,omitempty"`

	// The ID of a managed prefix list destination of the route.
	DestinationPrefixListId string `json:"destinationPrefixListId,omitempty" yaml:"destinationPrefixListId,omitempty"`

	// VPC Endpoint ID.
	VpcEndpointId string `json:"vpcEndpointId,omitempty" yaml:"vpcEndpointId,omitempty"`

	// VPC Peering ID.
	VpcPeeringConnectionId string `json:"vpcPeeringConnectionId,omitempty" yaml:"vpcPeeringConnectionId,omitempty"`

	// CIDR block of the route.
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`

	// EC2 instance ID.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// IPv6 CIDR block of the route.
	Ipv6CidrBlock string `json:"ipv6CidrBlock,omitempty" yaml:"ipv6CidrBlock,omitempty"`

	// ID of the Carrier Gateway.
	CarrierGatewayId string `json:"carrierGatewayId,omitempty" yaml:"carrierGatewayId,omitempty"`

	// ID of an Internet Gateway or Virtual Private Gateway which is connected to the Route Table (not exported if not passed as a parameter).
	GatewayId string `json:"gatewayId,omitempty" yaml:"gatewayId,omitempty"`

	// ID of the elastic network interface (eni) to use.
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`

	// EC2 Transit Gateway ID.
	TransitGatewayId string `json:"transitGatewayId,omitempty" yaml:"transitGatewayId,omitempty"`
}
