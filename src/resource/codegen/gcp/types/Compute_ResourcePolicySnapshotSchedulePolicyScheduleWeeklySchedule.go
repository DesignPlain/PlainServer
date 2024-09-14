package types

type Compute_ResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule struct {
	/*
	   May contain up to seven (one for each day of the week) snapshot times.
	   Structure is documented below.
	*/
	DayOfWeeks []Compute_ResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeek `json:"dayOfWeeks,omitempty" yaml:"dayOfWeeks,omitempty"`
}
