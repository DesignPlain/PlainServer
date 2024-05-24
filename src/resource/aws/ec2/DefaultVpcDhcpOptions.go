package ec2

type DefaultVpcDhcpOptions struct {
	// The ID of the AWS account that owns the DHCP options set.
	OwnerId string `json:"ownerId,omitempty" yaml:"ownerId,omitempty"`

	// A map of tags to assign to the resource.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
