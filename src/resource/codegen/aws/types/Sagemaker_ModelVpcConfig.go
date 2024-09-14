package types

type Sagemaker_ModelVpcConfig struct {
	//
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	//
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`
}
