package ec2

type LocalGatewayRoute struct {
	// IPv4 CIDR range used for destination matches. Routing decisions are based on the most specific match.
	DestinationCidrBlock string `json:"destinationCidrBlock,omitempty" yaml:"destinationCidrBlock,omitempty"`

	// Identifier of EC2 Local Gateway Route Table.
	LocalGatewayRouteTableId string `json:"localGatewayRouteTableId,omitempty" yaml:"localGatewayRouteTableId,omitempty"`

	// Identifier of EC2 Local Gateway Virtual Interface Group.
	LocalGatewayVirtualInterfaceGroupId string `json:"localGatewayVirtualInterfaceGroupId,omitempty" yaml:"localGatewayVirtualInterfaceGroupId,omitempty"`
}
