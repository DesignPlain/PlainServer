package types

type Opensearchingest_PipelineVpcOptions struct {
	// A list of security groups associated with the VPC endpoint.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// A list of subnet IDs associated with the VPC endpoint.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
}
