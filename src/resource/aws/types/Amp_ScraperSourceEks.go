package types

type Amp_ScraperSourceEks struct {
	//
	ClusterArn string `json:"clusterArn,omitempty" yaml:"clusterArn,omitempty"`

	// List of the security group IDs for the Amazon EKS cluster VPC configuration.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// List of subnet IDs. Must be in at least two different availability zones.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
}
