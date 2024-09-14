package types

type Looker_InstanceMaintenanceWindowStartTime struct {
	// Minutes of hour of day. Must be from 0 to 59.
	Minutes int `json:"minutes,omitempty" yaml:"minutes,omitempty"`

	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	Nanos int `json:"nanos,omitempty" yaml:"nanos,omitempty"`

	// Seconds of minutes of the time. Must normally be from 0 to 59.
	Seconds int `json:"seconds,omitempty" yaml:"seconds,omitempty"`

	// Hours of day in 24 hour format. Should be from 0 to 23.
	Hours int `json:"hours,omitempty" yaml:"hours,omitempty"`
}
