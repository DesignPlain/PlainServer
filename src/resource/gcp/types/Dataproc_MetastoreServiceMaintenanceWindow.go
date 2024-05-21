package types

type Dataproc_MetastoreServiceMaintenanceWindow struct {
	/*
	   The day of week, when the window starts.
	   Possible values are: `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, `SUNDAY`.
	*/
	DayOfWeek string `json:"dayOfWeek,omitempty" yaml:"dayOfWeek,omitempty"`

	// The hour of day (0-23) when the window starts.
	HourOfDay int `json:"hourOfDay,omitempty" yaml:"hourOfDay,omitempty"`
}
