package types

type Ecs_ServiceCapacityProviderStrategy struct {
	// Number of tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined.
	Base int `json:"base,omitempty" yaml:"base,omitempty"`

	// Short name of the capacity provider.
	CapacityProvider string `json:"capacityProvider,omitempty" yaml:"capacityProvider,omitempty"`

	// Relative percentage of the total number of launched tasks that should use the specified capacity provider.
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}
