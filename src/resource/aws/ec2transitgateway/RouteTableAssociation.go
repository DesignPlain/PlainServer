package ec2transitgateway

type RouteTableAssociation struct {
	// Boolean whether the Gateway Attachment should remove any current Route Table association before associating with the specified Route Table. Default value: `false`. This argument is intended for use with EC2 Transit Gateways shared into the current account, otherwise the `transit_gateway_default_route_table_association` argument of the `aws.ec2transitgateway.VpcAttachment` resource should be used.
	ReplaceExistingAssociation bool `json:"replaceExistingAssociation,omitempty" yaml:"replaceExistingAssociation,omitempty"`

	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId string `json:"transitGatewayAttachmentId,omitempty" yaml:"transitGatewayAttachmentId,omitempty"`

	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId string `json:"transitGatewayRouteTableId,omitempty" yaml:"transitGatewayRouteTableId,omitempty"`
}
