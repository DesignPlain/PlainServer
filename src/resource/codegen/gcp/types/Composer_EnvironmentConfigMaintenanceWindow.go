package types

type Composer_EnvironmentConfigMaintenanceWindow struct {
	// Maintenance window recurrence. Format is a subset of RFC-5545 (https://tools.ietf.org/html/rfc5545) 'RRULE'. The only allowed values for 'FREQ' field are 'FREQ=DAILY' and 'FREQ=WEEKLY;BYDAY=...'. Example values: 'FREQ=WEEKLY;BYDAY=TU,WE', 'FREQ=DAILY'.
	Recurrence string `json:"recurrence,omitempty" yaml:"recurrence,omitempty"`

	// Start time of the first recurrence of the maintenance window.
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`

	// Maintenance window end time. It is used only to calculate the duration of the maintenance window. The value for end-time must be in the future, relative to 'start_time'.
	EndTime string `json:"endTime,omitempty" yaml:"endTime,omitempty"`
}
