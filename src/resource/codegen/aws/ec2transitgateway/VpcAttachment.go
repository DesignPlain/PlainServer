package ec2transitgateway

type VpcAttachment struct {
	// Whether IPv6 support is enabled. Valid values: `disable`, `enable`. Default value: `disable`.
	Ipv6Support string `json:"ipv6Support,omitempty" yaml:"ipv6Support,omitempty"`

	// Key-value tags for the EC2 Transit Gateway VPC Attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTableAssociation bool `json:"transitGatewayDefaultRouteTableAssociation,omitempty" yaml:"transitGatewayDefaultRouteTableAssociation,omitempty"`

	// Identifier of EC2 Transit Gateway.
	TransitGatewayId string `json:"transitGatewayId,omitempty" yaml:"transitGatewayId,omitempty"`

	// Identifier of EC2 VPC.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// Whether Appliance Mode support is enabled. If enabled, a traffic flow between a source and destination uses the same Availability Zone for the VPC attachment for the lifetime of that flow. Valid values: `disable`, `enable`. Default value: `disable`.
	ApplianceModeSupport string `json:"applianceModeSupport,omitempty" yaml:"applianceModeSupport,omitempty"`

	// Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	DnsSupport string `json:"dnsSupport,omitempty" yaml:"dnsSupport,omitempty"`

	// Identifiers of EC2 Subnets.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTablePropagation bool `json:"transitGatewayDefaultRouteTablePropagation,omitempty" yaml:"transitGatewayDefaultRouteTablePropagation,omitempty"`
}
