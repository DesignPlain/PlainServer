package types

type Pipes_PipeSourceParametersSelfManagedKafkaParametersVpc struct {
	//
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	//
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`
}
