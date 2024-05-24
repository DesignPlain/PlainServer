package ec2

type VpnGateway struct {
	// The Autonomous System Number (ASN) for the Amazon side of the gateway. If you don't specify an ASN, the virtual private gateway is created with the default ASN.
	AmazonSideAsn string `json:"amazonSideAsn,omitempty" yaml:"amazonSideAsn,omitempty"`

	// The Availability Zone for the virtual private gateway.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The VPC ID to create in.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
