package types

type Connect_HoursOfOperationConfigStartTime struct {
	// Specifies the hour of opening.
	Hours int `json:"hours,omitempty" yaml:"hours,omitempty"`

	// Specifies the minute of opening.
	Minutes int `json:"minutes,omitempty" yaml:"minutes,omitempty"`
}
