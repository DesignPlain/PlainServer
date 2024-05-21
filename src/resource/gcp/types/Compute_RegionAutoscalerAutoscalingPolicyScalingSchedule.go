package types

type Compute_RegionAutoscalerAutoscalingPolicyScalingSchedule struct {
	// The start timestamps of time intervals when this scaling schedule should provide a scaling signal. This field uses the extended cron format (with an optional year field).
	Schedule string `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// The time zone to be used when interpreting the schedule. The value of this field must be a time zone name from the tz database: http://en.wikipedia.org/wiki/Tz_database.
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A boolean value that specifies if a scaling schedule can influence autoscaler recommendations. If set to true, then a scaling schedule has no effect.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	// The duration of time intervals (in seconds) for which this scaling schedule will be running. The minimum allowed value is 300.
	DurationSec int `json:"durationSec,omitempty" yaml:"durationSec,omitempty"`

	// Minimum number of VM instances that autoscaler will recommend in time intervals starting according to schedule.
	MinRequiredReplicas int `json:"minRequiredReplicas,omitempty" yaml:"minRequiredReplicas,omitempty"`

	// The identifier for this object. Format specified above.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
