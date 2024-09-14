package types

type Cloudwatch_EventTargetEcsTargetCapacityProviderStrategy struct {
	// Short name of the capacity provider.
	CapacityProvider string `json:"capacityProvider,omitempty" yaml:"capacityProvider,omitempty"`

	// The weight value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider. The weight value is taken into consideration after the base value, if defined, is satisfied.
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`

	// The base value designates how many tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined. Defaults to `0`.
	Base int `json:"base,omitempty" yaml:"base,omitempty"`
}
