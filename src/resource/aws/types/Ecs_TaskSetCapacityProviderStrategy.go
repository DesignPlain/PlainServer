package types

type Ecs_TaskSetCapacityProviderStrategy struct {
	// The relative percentage of the total number of launched tasks that should use the specified capacity provider.
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`

	// The number of tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined.
	Base int `json:"base,omitempty" yaml:"base,omitempty"`

	// The short name or full Amazon Resource Name (ARN) of the capacity provider.
	CapacityProvider string `json:"capacityProvider,omitempty" yaml:"capacityProvider,omitempty"`
}
