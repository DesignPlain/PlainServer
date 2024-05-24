package types

type Lambda_FunctionVpcConfig struct {
	// Allows outbound IPv6 traffic on VPC functions that are connected to dual-stack subnets. Default is `false`.
	Ipv6AllowedForDualStack bool `json:"ipv6AllowedForDualStack,omitempty" yaml:"ipv6AllowedForDualStack,omitempty"`

	// List of security group IDs associated with the Lambda function.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// List of subnet IDs associated with the Lambda function.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	//
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
