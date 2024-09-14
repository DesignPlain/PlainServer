package types

type Compute_getResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule struct {
	// The number of hours between snapshots.
	HoursInCycle int `json:"hoursInCycle,omitempty" yaml:"hoursInCycle,omitempty"`

	/*
	   Time within the window to start the operations.
	   It must be in an hourly format "HH:MM",
	   where HH : [00-23] and MM : [00] GMT.
	   eg: 21:00
	*/
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`
}
