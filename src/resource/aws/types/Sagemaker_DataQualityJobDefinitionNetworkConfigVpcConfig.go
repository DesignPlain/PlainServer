package types

type Sagemaker_DataQualityJobDefinitionNetworkConfigVpcConfig struct {
	// The VPC security group IDs, in the form sg-xxxxxxxx. Specify the security groups for the VPC that is specified in the `subnets` field.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The ID of the subnets in the VPC to which you want to connect your training job or model.
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`
}
