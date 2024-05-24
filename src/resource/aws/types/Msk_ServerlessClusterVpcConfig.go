package types

type Msk_ServerlessClusterVpcConfig struct {
	// Specifies up to five security groups that control inbound and outbound traffic for the serverless cluster.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// A list of subnets in at least two different Availability Zones that host your client applications.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
}
