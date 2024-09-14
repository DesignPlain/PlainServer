package types

type Codebuild_ProjectVpcConfig struct {
	// Security group IDs to assign to running builds.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// Subnet IDs within which to run builds.
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`

	// ID of the VPC within which to run builds.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
