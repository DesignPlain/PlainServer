package ec2

import types "libds/aws/types"

type DefaultRouteTable struct {
	// Set of objects. Detailed below
	Routes []types.Ec2_DefaultRouteTableRoute `json:"routes,omitempty" yaml:"routes,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   ID of the default route table.

	   The following arguments are optional:
	*/
	DefaultRouteTableId string `json:"defaultRouteTableId,omitempty" yaml:"defaultRouteTableId,omitempty"`

	// List of virtual gateways for propagation.
	PropagatingVgws []string `json:"propagatingVgws,omitempty" yaml:"propagatingVgws,omitempty"`
}
