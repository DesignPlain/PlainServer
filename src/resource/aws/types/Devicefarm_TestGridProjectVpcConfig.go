package types

type Devicefarm_TestGridProjectVpcConfig struct {
	// A list of VPC security group IDs in your Amazon VPC.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// A list of VPC subnet IDs in your Amazon VPC.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// The ID of the Amazon VPC.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
