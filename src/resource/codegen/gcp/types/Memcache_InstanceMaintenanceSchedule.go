package types

type Memcache_InstanceMaintenanceSchedule struct {
	/*
	   (Output)
	   Output only. The deadline that the maintenance schedule start time
	   can not go beyond, including reschedule.
	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	   resolution and up to nine fractional digits.
	*/
	ScheduleDeadlineTime string `json:"scheduleDeadlineTime,omitempty" yaml:"scheduleDeadlineTime,omitempty"`

	/*
	   Required. Start time of the window in UTC time.
	   Structure is documented below.
	*/
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`

	/*
	   (Output)
	   Output only. The end time of any upcoming scheduled maintenance for this instance.
	   A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	   resolution and up to nine fractional digits.
	*/
	EndTime string `json:"endTime,omitempty" yaml:"endTime,omitempty"`
}
