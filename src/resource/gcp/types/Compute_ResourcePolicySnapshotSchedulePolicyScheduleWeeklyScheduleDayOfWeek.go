package types

type Compute_ResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeek struct {
	/*
	   The day of the week to create the snapshot. e.g. MONDAY
	   Possible values are: `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, `SUNDAY`.
	*/
	Day string `json:"day,omitempty" yaml:"day,omitempty"`

	/*
	   Time within the window to start the operations.
	   It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.
	*/
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`
}
