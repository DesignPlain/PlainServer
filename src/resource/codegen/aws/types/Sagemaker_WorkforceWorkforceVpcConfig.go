package types

type Sagemaker_WorkforceWorkforceVpcConfig struct {
	// The ID of the VPC that the workforce uses for communication.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// The VPC security group IDs. The security groups must be for the same VPC as specified in the subnet.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The ID of the subnets in the VPC that you want to connect.
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`

	// The IDs for the VPC service endpoints of your VPC workforce.
	VpcEndpointId string `json:"vpcEndpointId,omitempty" yaml:"vpcEndpointId,omitempty"`
}
