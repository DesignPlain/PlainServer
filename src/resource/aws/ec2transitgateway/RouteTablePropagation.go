package ec2transitgateway

type RouteTablePropagation struct {
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId string `json:"transitGatewayAttachmentId,omitempty" yaml:"transitGatewayAttachmentId,omitempty"`

	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId string `json:"transitGatewayRouteTableId,omitempty" yaml:"transitGatewayRouteTableId,omitempty"`
}
