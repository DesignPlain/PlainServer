package autoscaling

type Schedule struct {
	// The maximum size of the Auto Scaling group. Set to `-1` if you don't want to change the maximum size at the scheduled time. Defaults to `0`.
	MaxSize int `json:"maxSize,omitempty" yaml:"maxSize,omitempty"`

	// The minimum size of the Auto Scaling group. Set to `-1` if you don't want to change the minimum size at the scheduled time. Defaults to `0`.
	MinSize int `json:"minSize,omitempty" yaml:"minSize,omitempty"`

	// The recurring schedule for this action specified using the Unix cron syntax format.
	Recurrence string `json:"recurrence,omitempty" yaml:"recurrence,omitempty"`

	// The date and time for the recurring schedule to start, in UTC with the format `"YYYY-MM-DDThh:mm:ssZ"` (e.g. `"2021-06-01T00:00:00Z"`).
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`

	// The name of the Auto Scaling group.
	AutoscalingGroupName string `json:"autoscalingGroupName,omitempty" yaml:"autoscalingGroupName,omitempty"`

	// The initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain. Set to `-1` if you don't want to change the desired capacity at the scheduled time. Defaults to `0`.
	DesiredCapacity int `json:"desiredCapacity,omitempty" yaml:"desiredCapacity,omitempty"`

	/*
	   Specifies the time zone for a cron expression. Valid values are the canonical names of the IANA time zones (such as `Etc/GMT+9` or `Pacific/Tahiti`).

	   > --NOTE:-- When `start_time` and `end_time` are specified with `recurrence` , they form the boundaries of when the recurring action will start and stop.
	*/
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`

	// The date and time for the recurring schedule to end, in UTC with the format `"YYYY-MM-DDThh:mm:ssZ"` (e.g. `"2021-06-01T00:00:00Z"`).
	EndTime string `json:"endTime,omitempty" yaml:"endTime,omitempty"`

	/*
	   The name of this scaling action.

	   The following arguments are optional:
	*/
	ScheduledActionName string `json:"scheduledActionName,omitempty" yaml:"scheduledActionName,omitempty"`
}
