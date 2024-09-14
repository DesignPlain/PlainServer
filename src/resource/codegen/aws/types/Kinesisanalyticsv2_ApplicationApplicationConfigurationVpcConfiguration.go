package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationVpcConfiguration struct {
	// The Security Group IDs used by the VPC configuration.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The Subnet IDs used by the VPC configuration.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	//
	VpcConfigurationId string `json:"vpcConfigurationId,omitempty" yaml:"vpcConfigurationId,omitempty"`

	//
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
