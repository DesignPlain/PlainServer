package types

type Bedrock_CustomModelVpcConfig struct {
	// VPC configuration security group IDs.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// VPC configuration subnets.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
}
