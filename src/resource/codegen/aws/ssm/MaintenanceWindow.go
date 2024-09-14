package ssm

type MaintenanceWindow struct {
	// The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
	ScheduleOffset int `json:"scheduleOffset,omitempty" yaml:"scheduleOffset,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A description for the maintenance window.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
	EndDate string `json:"endDate,omitempty" yaml:"endDate,omitempty"`

	// The schedule of the Maintenance Window in the form of a [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html).
	Schedule string `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// Whether the maintenance window is enabled. Default: `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The name of the maintenance window.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
	ScheduleTimezone string `json:"scheduleTimezone,omitempty" yaml:"scheduleTimezone,omitempty"`

	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
	StartDate string `json:"startDate,omitempty" yaml:"startDate,omitempty"`

	// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
	AllowUnassociatedTargets bool `json:"allowUnassociatedTargets,omitempty" yaml:"allowUnassociatedTargets,omitempty"`

	// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
	Cutoff int `json:"cutoff,omitempty" yaml:"cutoff,omitempty"`

	// The duration of the Maintenance Window in hours.
	Duration int `json:"duration,omitempty" yaml:"duration,omitempty"`
}
