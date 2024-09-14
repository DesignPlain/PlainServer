package types

type Ecs_getTaskExecutionCapacityProviderStrategy struct {
	// The number of tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined. Defaults to `0`.
	Base int `json:"base,omitempty" yaml:"base,omitempty"`

	// Name of the capacity provider.
	CapacityProvider string `json:"capacityProvider,omitempty" yaml:"capacityProvider,omitempty"`

	// The relative percentage of the total number of launched tasks that should use the specified capacity provider. The `weight` value is taken into consideration after the `base` count of tasks has been satisfied. Defaults to `0`.
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}
