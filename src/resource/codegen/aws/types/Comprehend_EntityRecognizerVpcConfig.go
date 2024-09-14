package types

type Comprehend_EntityRecognizerVpcConfig struct {
	// List of security group IDs.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// List of VPC subnets.
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`
}
