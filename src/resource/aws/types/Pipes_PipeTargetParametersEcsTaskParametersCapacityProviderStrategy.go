package types

type Pipes_PipeTargetParametersEcsTaskParametersCapacityProviderStrategy struct {
	// The base value designates how many tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined. If no value is specified, the default value of 0 is used. Maximum value of 100,000.
	Base int `json:"base,omitempty" yaml:"base,omitempty"`

	// The short name of the capacity provider. Maximum value of 255.
	CapacityProvider string `json:"capacityProvider,omitempty" yaml:"capacityProvider,omitempty"`

	// The weight value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider. The weight value is taken into consideration after the base value, if defined, is satisfied. Maximum value of 1,000.
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}
