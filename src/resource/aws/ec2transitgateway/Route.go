package ec2transitgateway

type Route struct {
	// Identifier of EC2 Transit Gateway Attachment (required if `blackhole` is set to false).
	TransitGatewayAttachmentId string `json:"transitGatewayAttachmentId,omitempty" yaml:"transitGatewayAttachmentId,omitempty"`

	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId string `json:"transitGatewayRouteTableId,omitempty" yaml:"transitGatewayRouteTableId,omitempty"`

	// Indicates whether to drop traffic that matches this route (default to `false`).
	Blackhole bool `json:"blackhole,omitempty" yaml:"blackhole,omitempty"`

	// IPv4 or IPv6 RFC1924 CIDR used for destination matches. Routing decisions are based on the most specific match.
	DestinationCidrBlock string `json:"destinationCidrBlock,omitempty" yaml:"destinationCidrBlock,omitempty"`
}
