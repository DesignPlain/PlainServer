package types

type Compute_ResourcePolicyInstanceSchedulePolicy struct {
	/*
	   Specifies the schedule for starting instances.
	   Structure is documented below.
	*/
	VmStartSchedule Compute_ResourcePolicyInstanceSchedulePolicyVmStartSchedule `json:"vmStartSchedule,omitempty" yaml:"vmStartSchedule,omitempty"`

	/*
	   Specifies the schedule for stopping instances.
	   Structure is documented below.
	*/
	VmStopSchedule Compute_ResourcePolicyInstanceSchedulePolicyVmStopSchedule `json:"vmStopSchedule,omitempty" yaml:"vmStopSchedule,omitempty"`

	// The expiration time of the schedule. The timestamp is an RFC3339 string.
	ExpirationTime string `json:"expirationTime,omitempty" yaml:"expirationTime,omitempty"`

	// The start time of the schedule. The timestamp is an RFC3339 string.
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`

	/*
	   Specifies the time zone to be used in interpreting the schedule. The value of this field must be a time zone name
	   from the tz database: http://en.wikipedia.org/wiki/Tz_database.
	*/
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`
}
