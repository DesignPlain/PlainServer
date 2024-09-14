package types

type Firestore_BackupScheduleWeeklyRecurrence struct {
	/*
	   The day of week to run.
	   Possible values are: `DAY_OF_WEEK_UNSPECIFIED`, `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, `SUNDAY`.
	*/
	Day string `json:"day,omitempty" yaml:"day,omitempty"`
}
