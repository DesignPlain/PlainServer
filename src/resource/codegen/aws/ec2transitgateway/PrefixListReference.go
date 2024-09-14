package ec2transitgateway

type PrefixListReference struct {
	// Indicates whether to drop traffic that matches the Prefix List. Defaults to `false`.
	Blackhole bool `json:"blackhole,omitempty" yaml:"blackhole,omitempty"`

	// Identifier of EC2 Prefix List.
	PrefixListId string `json:"prefixListId,omitempty" yaml:"prefixListId,omitempty"`

	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId string `json:"transitGatewayAttachmentId,omitempty" yaml:"transitGatewayAttachmentId,omitempty"`

	/*
	   Identifier of EC2 Transit Gateway Route Table.

	   The following arguments are optional:
	*/
	TransitGatewayRouteTableId string `json:"transitGatewayRouteTableId,omitempty" yaml:"transitGatewayRouteTableId,omitempty"`
}
