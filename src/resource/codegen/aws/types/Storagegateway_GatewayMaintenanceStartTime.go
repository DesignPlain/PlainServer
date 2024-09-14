package types

type Storagegateway_GatewayMaintenanceStartTime struct {
	// The day of the month component of the maintenance start time represented as an ordinal number from 1 to 28, where 1 represents the first day of the month and 28 represents the last day of the month.
	DayOfMonth string `json:"dayOfMonth,omitempty" yaml:"dayOfMonth,omitempty"`

	// The day of the week component of the maintenance start time week represented as an ordinal number from 0 to 6, where 0 represents Sunday and 6 Saturday.
	DayOfWeek string `json:"dayOfWeek,omitempty" yaml:"dayOfWeek,omitempty"`

	// The hour component of the maintenance start time represented as _hh_, where _hh_ is the hour (00 to 23). The hour of the day is in the time zone of the gateway.
	HourOfDay int `json:"hourOfDay,omitempty" yaml:"hourOfDay,omitempty"`

	// The minute component of the maintenance start time represented as _mm_, where _mm_ is the minute (00 to 59). The minute of the hour is in the time zone of the gateway.
	MinuteOfHour int `json:"minuteOfHour,omitempty" yaml:"minuteOfHour,omitempty"`
}
