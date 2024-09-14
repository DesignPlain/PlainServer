package types

type Ssm_getContactsRotationRecurrenceWeeklySetting struct {
	//
	DayOfWeek string `json:"dayOfWeek,omitempty" yaml:"dayOfWeek,omitempty"`

	//
	HandOffTimes []Ssm_getContactsRotationRecurrenceWeeklySettingHandOffTime `json:"handOffTimes,omitempty" yaml:"handOffTimes,omitempty"`
}
