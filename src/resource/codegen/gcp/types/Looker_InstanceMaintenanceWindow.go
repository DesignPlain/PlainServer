package types

type Looker_InstanceMaintenanceWindow struct {
	/*
	   Required. Day of the week for this MaintenanceWindow (in UTC).
	   - MONDAY: Monday
	   - TUESDAY: Tuesday
	   - WEDNESDAY: Wednesday
	   - THURSDAY: Thursday
	   - FRIDAY: Friday
	   - SATURDAY: Saturday
	   - SUNDAY: Sunday
	   Possible values are: `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, `SUNDAY`.
	*/
	DayOfWeek string `json:"dayOfWeek,omitempty" yaml:"dayOfWeek,omitempty"`

	/*
	   Required. Start time of the window in UTC time.
	   Structure is documented below.
	*/
	StartTime Looker_InstanceMaintenanceWindowStartTime `json:"startTime,omitempty" yaml:"startTime,omitempty"`
}
