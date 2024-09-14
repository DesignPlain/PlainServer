package types

type Apigee_EnvironmentNodeConfig struct {
	/*
	   The maximum total number of gateway nodes that the is reserved for all instances that
	   has the specified environment. If not specified, the default is determined by the
	   recommended maximum number of nodes for that gateway.
	*/
	MaxNodeCount string `json:"maxNodeCount,omitempty" yaml:"maxNodeCount,omitempty"`

	/*
	   The minimum total number of gateway nodes that the is reserved for all instances that
	   has the specified environment. If not specified, the default is determined by the
	   recommended minimum number of nodes for that gateway.
	*/
	MinNodeCount string `json:"minNodeCount,omitempty" yaml:"minNodeCount,omitempty"`

	/*
	   (Output)
	   The current total number of gateway nodes that each environment currently has across
	   all instances.
	*/
	CurrentAggregateNodeCount string `json:"currentAggregateNodeCount,omitempty" yaml:"currentAggregateNodeCount,omitempty"`
}
