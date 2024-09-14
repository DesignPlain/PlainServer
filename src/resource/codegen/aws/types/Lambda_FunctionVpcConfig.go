package types

type Lambda_FunctionVpcConfig struct {
	// List of subnet IDs associated with the Lambda function.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// ID of the VPC.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// Allows outbound IPv6 traffic on VPC functions that are connected to dual-stack subnets. Default is `false`.
	Ipv6AllowedForDualStack bool `json:"ipv6AllowedForDualStack,omitempty" yaml:"ipv6AllowedForDualStack,omitempty"`

	// List of security group IDs associated with the Lambda function.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`
}
