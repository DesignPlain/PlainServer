package types

type Iot_TopicRuleDestinationVpcConfiguration struct {
	// The security groups of the VPC destination.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	// The subnet IDs of the VPC destination.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// The ID of the VPC.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// The ARN of a role that has permission to create and attach to elastic network interfaces (ENIs).
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
