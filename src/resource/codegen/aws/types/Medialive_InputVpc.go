package types

type Medialive_InputVpc struct {
	// A list of up to 5 EC2 VPC security group IDs to attach to the Input.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// A list of 2 VPC subnet IDs from the same VPC.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
}
