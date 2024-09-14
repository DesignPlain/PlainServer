package types

type Connect_getHoursOfOperationConfig struct {
	// Day that the hours of operation applies to.
	Day string `json:"day,omitempty" yaml:"day,omitempty"`

	// End time block specifies the time that your contact center closes. The `end_time` is documented below.
	EndTimes []Connect_getHoursOfOperationConfigEndTime `json:"endTimes,omitempty" yaml:"endTimes,omitempty"`

	// Start time block specifies the time that your contact center opens. The `start_time` is documented below.
	StartTimes []Connect_getHoursOfOperationConfigStartTime `json:"startTimes,omitempty" yaml:"startTimes,omitempty"`
}
