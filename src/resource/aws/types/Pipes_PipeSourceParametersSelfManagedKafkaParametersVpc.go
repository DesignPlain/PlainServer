package types

type Pipes_PipeSourceParametersSelfManagedKafkaParametersVpc struct {
	// List of security groups associated with the stream. These security groups must all be in the same VPC. You can specify as many as five security groups. If you do not specify a security group, the default security group for the VPC is used.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	// List of the subnets associated with the stream. These subnets must all be in the same VPC. You can specify as many as 16 subnets.
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`
}
