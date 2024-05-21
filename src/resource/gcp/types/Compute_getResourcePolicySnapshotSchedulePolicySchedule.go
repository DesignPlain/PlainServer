package types

type Compute_getResourcePolicySnapshotSchedulePolicySchedule struct {
	// The policy will execute every nth day at the specified time.
	DailySchedules []Compute_getResourcePolicySnapshotSchedulePolicyScheduleDailySchedule `json:"dailySchedules,omitempty" yaml:"dailySchedules,omitempty"`

	// The policy will execute every nth hour starting at the specified time.
	HourlySchedules []Compute_getResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule `json:"hourlySchedules,omitempty" yaml:"hourlySchedules,omitempty"`

	// Allows specifying a snapshot time for each day of the week.
	WeeklySchedules []Compute_getResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule `json:"weeklySchedules,omitempty" yaml:"weeklySchedules,omitempty"`
}
