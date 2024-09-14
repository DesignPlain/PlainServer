package types

type Scheduler_ScheduleTargetEcsParametersPlacementStrategy struct {
	// The field to apply the placement strategy against.
	Field string `json:"field,omitempty" yaml:"field,omitempty"`

	// The type of placement strategy. One of: `random`, `spread`, `binpack`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
