package types

type Quicksight_RefreshScheduleScheduleScheduleFrequencyRefreshOnDay struct {
	// The day of the month that you want to schedule refresh on.
	DayOfMonth string `json:"dayOfMonth,omitempty" yaml:"dayOfMonth,omitempty"`

	// The day of the week that you want to schedule a refresh on. Valid values are `SUNDAY`, `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY` and `SATURDAY`.
	DayOfWeek string `json:"dayOfWeek,omitempty" yaml:"dayOfWeek,omitempty"`
}
