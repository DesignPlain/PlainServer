package types

type Compute_ResourcePolicySnapshotSchedulePolicySchedule struct {
	/*
	   The policy will execute every nth day at the specified time.
	   Structure is documented below.
	*/
	DailySchedule Compute_ResourcePolicySnapshotSchedulePolicyScheduleDailySchedule `json:"dailySchedule,omitempty" yaml:"dailySchedule,omitempty"`

	/*
	   The policy will execute every nth hour starting at the specified time.
	   Structure is documented below.
	*/
	HourlySchedule Compute_ResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule `json:"hourlySchedule,omitempty" yaml:"hourlySchedule,omitempty"`

	/*
	   Allows specifying a snapshot time for each day of the week.
	   Structure is documented below.
	*/
	WeeklySchedule Compute_ResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule `json:"weeklySchedule,omitempty" yaml:"weeklySchedule,omitempty"`
}
