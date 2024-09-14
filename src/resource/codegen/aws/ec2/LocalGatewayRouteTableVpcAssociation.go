package ec2

type LocalGatewayRouteTableVpcAssociation struct {
	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Identifier of EC2 VPC.

	   The following arguments are optional:
	*/
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// Identifier of EC2 Local Gateway Route Table.
	LocalGatewayRouteTableId string `json:"localGatewayRouteTableId,omitempty" yaml:"localGatewayRouteTableId,omitempty"`
}
