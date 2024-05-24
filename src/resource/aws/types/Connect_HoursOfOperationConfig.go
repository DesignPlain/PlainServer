package types

type Connect_HoursOfOperationConfig struct {
	// Specifies the day that the hours of operation applies to.
	Day string `json:"day,omitempty" yaml:"day,omitempty"`

	// A end time block specifies the time that your contact center closes. The `end_time` is documented below.
	EndTime Connect_HoursOfOperationConfigEndTime `json:"endTime,omitempty" yaml:"endTime,omitempty"`

	// A start time block specifies the time that your contact center opens. The `start_time` is documented below.
	StartTime Connect_HoursOfOperationConfigStartTime `json:"startTime,omitempty" yaml:"startTime,omitempty"`
}
