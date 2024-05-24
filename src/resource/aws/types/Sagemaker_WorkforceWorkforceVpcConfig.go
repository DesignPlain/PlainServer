package types

type Sagemaker_WorkforceWorkforceVpcConfig struct {
	// The VPC security group IDs. The security groups must be for the same VPC as specified in the subnet.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The ID of the subnets in the VPC that you want to connect.
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`

	//
	VpcEndpointId string `json:"vpcEndpointId,omitempty" yaml:"vpcEndpointId,omitempty"`

	// The ID of the VPC that the workforce uses for communication.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
