package types

type Compute_SecurityScanConfigSchedule struct {
	// The duration of time between executions in days
	IntervalDurationDays int `json:"intervalDurationDays,omitempty" yaml:"intervalDurationDays,omitempty"`

	/*
	   A timestamp indicates when the next run will be scheduled. The value is refreshed
	   by the server after each run. If unspecified, it will default to current server time,
	   which means the scan will be scheduled to start immediately.
	*/
	ScheduleTime string `json:"scheduleTime,omitempty" yaml:"scheduleTime,omitempty"`
}
