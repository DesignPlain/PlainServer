package types

type Ecs_TaskSetScale struct {
	// The unit of measure for the scale value. Default: `PERCENT`.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	// The value, specified as a percent total of a service's `desiredCount`, to scale the task set. Defaults to `0` if not specified. Accepted values are numbers between 0.0 and 100.0.
	Value float64 `json:"value,omitempty" yaml:"value,omitempty"`
}
