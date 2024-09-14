package types

type Osconfig_PatchDeploymentRecurringScheduleWeekly struct {
	/*
	   IANA Time Zone Database time zone, e.g. "America/New_York".
	   Possible values are: `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, `SUNDAY`.
	*/
	DayOfWeek string `json:"dayOfWeek,omitempty" yaml:"dayOfWeek,omitempty"`
}
