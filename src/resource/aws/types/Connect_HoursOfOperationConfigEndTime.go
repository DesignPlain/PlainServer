package types

type Connect_HoursOfOperationConfigEndTime struct {
	// Specifies the hour of closing.
	Hours int `json:"hours,omitempty" yaml:"hours,omitempty"`

	// Specifies the minute of closing.
	Minutes int `json:"minutes,omitempty" yaml:"minutes,omitempty"`
}
