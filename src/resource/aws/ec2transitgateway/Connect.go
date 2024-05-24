package ec2transitgateway

type Connect struct {
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId string `json:"transitGatewayId,omitempty" yaml:"transitGatewayId,omitempty"`

	// The underlaying VPC attachment
	TransportAttachmentId string `json:"transportAttachmentId,omitempty" yaml:"transportAttachmentId,omitempty"`

	// The tunnel protocol. Valid values: `gre`. Default is `gre`.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// Key-value tags for the EC2 Transit Gateway Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTableAssociation bool `json:"transitGatewayDefaultRouteTableAssociation,omitempty" yaml:"transitGatewayDefaultRouteTableAssociation,omitempty"`

	// Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTablePropagation bool `json:"transitGatewayDefaultRouteTablePropagation,omitempty" yaml:"transitGatewayDefaultRouteTablePropagation,omitempty"`
}
