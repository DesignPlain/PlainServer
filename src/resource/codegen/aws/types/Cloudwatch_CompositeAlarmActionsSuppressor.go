package types

type Cloudwatch_CompositeAlarmActionsSuppressor struct {
	// Can be an AlarmName or an Amazon Resource Name (ARN) from an existing alarm.
	Alarm string `json:"alarm,omitempty" yaml:"alarm,omitempty"`

	// The maximum time in seconds that the composite alarm waits after suppressor alarm goes out of the `ALARM` state. After this time, the composite alarm performs its actions.
	ExtensionPeriod int `json:"extensionPeriod,omitempty" yaml:"extensionPeriod,omitempty"`

	// The maximum time in seconds that the composite alarm waits for the suppressor alarm to go into the `ALARM` state. After this time, the composite alarm performs its actions.
	WaitPeriod int `json:"waitPeriod,omitempty" yaml:"waitPeriod,omitempty"`
}
