package devicefarm

import types "libds/aws/types"

type TestGridProject struct {
	// The VPC security groups and subnets that are attached to a project. See VPC Config below.
	VpcConfig types.Devicefarm_TestGridProjectVpcConfig `json:"vpcConfig,omitempty" yaml:"vpcConfig,omitempty"`

	// Human-readable description of the project.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the Selenium testing project.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
