package ec2

type CarrierGateway struct {
	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of the VPC to associate with the carrier gateway.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
