package types

type Mq_BrokerMaintenanceWindowStartTime struct {
	// Day of the week, e.g., `MONDAY`, `TUESDAY`, or `WEDNESDAY`.
	DayOfWeek string `json:"dayOfWeek,omitempty" yaml:"dayOfWeek,omitempty"`

	// Time, in 24-hour format, e.g., `02:00`.
	TimeOfDay string `json:"timeOfDay,omitempty" yaml:"timeOfDay,omitempty"`

	// Time zone in either the Country/City format or the UTC offset format, e.g., `CET`.
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`
}
