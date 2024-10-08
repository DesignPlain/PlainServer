package types

type Compute_ResourcePolicySnapshotSchedulePolicyScheduleDailySchedule struct {
	// Defines a schedule with units measured in days. The value determines how many days pass between the start of each cycle. Days in cycle for snapshot schedule policy must be 1.
	DaysInCycle int `json:"daysInCycle,omitempty" yaml:"daysInCycle,omitempty"`

	/*
	   This must be in UTC format that resolves to one of
	   00:00, 04:00, 08:00, 12:00, 16:00, or 20:00. For example,
	   both 13:00-5 and 08:00 are valid.
	*/
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`
}
