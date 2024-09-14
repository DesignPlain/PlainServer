package types

type Redis_getInstanceMaintenanceSchedule struct {
	/*
	   Output only. The deadline that the maintenance schedule start time
	   can not go beyond, including reschedule.
	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	   resolution and up to nine fractional digits.
	*/
	ScheduleDeadlineTime string `json:"scheduleDeadlineTime,omitempty" yaml:"scheduleDeadlineTime,omitempty"`

	/*
	   Output only. The start time of any upcoming scheduled maintenance for this instance.
	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	   resolution and up to nine fractional digits.
	*/
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`

	/*
	   Output only. The end time of any upcoming scheduled maintenance for this instance.
	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	   resolution and up to nine fractional digits.
	*/
	EndTime string `json:"endTime,omitempty" yaml:"endTime,omitempty"`
}
