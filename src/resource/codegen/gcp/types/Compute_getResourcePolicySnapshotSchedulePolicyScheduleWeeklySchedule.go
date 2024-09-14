package types

type Compute_getResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule struct {
	// May contain up to seven (one for each day of the week) snapshot times.
	DayOfWeeks []Compute_getResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeek `json:"dayOfWeeks,omitempty" yaml:"dayOfWeeks,omitempty"`
}
