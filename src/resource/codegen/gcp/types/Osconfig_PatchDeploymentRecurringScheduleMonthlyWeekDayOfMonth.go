package types

type Osconfig_PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth struct {
	// Week number in a month. 1-4 indicates the 1st to 4th week of the month. -1 indicates the last week of the month.
	WeekOrdinal int `json:"weekOrdinal,omitempty" yaml:"weekOrdinal,omitempty"`

	/*
	   A day of the week.
	   Possible values are: `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, `SUNDAY`.
	*/
	DayOfWeek string `json:"dayOfWeek,omitempty" yaml:"dayOfWeek,omitempty"`

	// Represents the number of days before or after the given week day of month that the patch deployment is scheduled for.
	DayOffset int `json:"dayOffset,omitempty" yaml:"dayOffset,omitempty"`
}
