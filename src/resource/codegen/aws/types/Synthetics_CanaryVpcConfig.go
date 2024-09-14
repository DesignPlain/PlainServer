package types

type Synthetics_CanaryVpcConfig struct {
	// IDs of the security groups for this canary.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// IDs of the subnets where this canary is to run.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// ID of the VPC where this canary is to run.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
