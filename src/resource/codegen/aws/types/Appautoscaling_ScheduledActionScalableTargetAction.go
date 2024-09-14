package types

type Appautoscaling_ScheduledActionScalableTargetAction struct {
	// Maximum capacity. At least one of `max_capacity` or `min_capacity` must be set.
	MaxCapacity int `json:"maxCapacity,omitempty" yaml:"maxCapacity,omitempty"`

	// Minimum capacity. At least one of `min_capacity` or `max_capacity` must be set.
	MinCapacity int `json:"minCapacity,omitempty" yaml:"minCapacity,omitempty"`
}
