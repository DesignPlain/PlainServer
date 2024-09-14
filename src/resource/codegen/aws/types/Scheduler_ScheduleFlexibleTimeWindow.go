package types

type Scheduler_ScheduleFlexibleTimeWindow struct {
	// Maximum time window during which a schedule can be invoked. Ranges from `1` to `1440` minutes.
	MaximumWindowInMinutes int `json:"maximumWindowInMinutes,omitempty" yaml:"maximumWindowInMinutes,omitempty"`

	// Determines whether the schedule is invoked within a flexible time window. One of: `OFF`, `FLEXIBLE`.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`
}
