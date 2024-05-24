package types

type Grafana_WorkspaceVpcConfiguration struct {
	// The list of Amazon EC2 security group IDs attached to the Amazon VPC for your Grafana workspace to connect.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The list of Amazon EC2 subnet IDs created in the Amazon VPC for your Grafana workspace to connect.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
}
