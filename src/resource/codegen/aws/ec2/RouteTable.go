package ec2

import types "libds/aws/types"

type RouteTable struct {
	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The VPC ID.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// A list of virtual gateways for propagation.
	PropagatingVgws []string `json:"propagatingVgws,omitempty" yaml:"propagatingVgws,omitempty"`

	/*
	   A list of route objects. Their keys are documented below.
	   This means that omitting this argument is interpreted as ignoring any existing routes. To remove all managed routes an empty list should be specified. See the example above.
	*/
	Routes []types.Ec2_RouteTableRoute `json:"routes,omitempty" yaml:"routes,omitempty"`
}
