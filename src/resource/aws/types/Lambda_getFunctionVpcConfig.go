package types

type Lambda_getFunctionVpcConfig struct {
	//
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	//
	Ipv6AllowedForDualStack bool `json:"ipv6AllowedForDualStack,omitempty" yaml:"ipv6AllowedForDualStack,omitempty"`

	//
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	//
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
}
