package types

type Alloydb_ClusterAutomatedBackupPolicyWeeklyScheduleStartTime struct {
	// Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value "24:00:00" for scenarios like business closing time.
	Hours int `json:"hours,omitempty" yaml:"hours,omitempty"`

	// Minutes of hour of day. Currently, only the value 0 is supported.
	Minutes int `json:"minutes,omitempty" yaml:"minutes,omitempty"`

	// Fractions of seconds in nanoseconds. Currently, only the value 0 is supported.
	Nanos int `json:"nanos,omitempty" yaml:"nanos,omitempty"`

	// Seconds of minutes of the time. Currently, only the value 0 is supported.
	Seconds int `json:"seconds,omitempty" yaml:"seconds,omitempty"`
}
