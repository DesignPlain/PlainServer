package types

type Scheduler_ScheduleTargetEcsParametersCapacityProviderStrategy struct {
	// How many tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined. Ranges from `0` (default) to `100000`.
	Base int `json:"base,omitempty" yaml:"base,omitempty"`

	// Short name of the capacity provider.
	CapacityProvider string `json:"capacityProvider,omitempty" yaml:"capacityProvider,omitempty"`

	// Designates the relative percentage of the total number of tasks launched that should use the specified capacity provider. The weight value is taken into consideration after the base value, if defined, is satisfied. Ranges from from `0` to `1000`.
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}
