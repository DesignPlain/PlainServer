package ec2transitgateway

type VpcAttachmentAccepter struct {
	// Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. Default value: `true`.
	TransitGatewayDefaultRouteTablePropagation bool `json:"transitGatewayDefaultRouteTablePropagation,omitempty" yaml:"transitGatewayDefaultRouteTablePropagation,omitempty"`

	// Key-value tags for the EC2 Transit Gateway VPC Attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of the EC2 Transit Gateway Attachment to manage.
	TransitGatewayAttachmentId string `json:"transitGatewayAttachmentId,omitempty" yaml:"transitGatewayAttachmentId,omitempty"`

	// Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. Default value: `true`.
	TransitGatewayDefaultRouteTableAssociation bool `json:"transitGatewayDefaultRouteTableAssociation,omitempty" yaml:"transitGatewayDefaultRouteTableAssociation,omitempty"`
}
