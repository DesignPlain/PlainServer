package types

type Ssm_getContactsRotationRecurrenceMonthlySetting struct {
	//
	DayOfMonth int `json:"dayOfMonth,omitempty" yaml:"dayOfMonth,omitempty"`

	//
	HandOffTimes []Ssm_getContactsRotationRecurrenceMonthlySettingHandOffTime `json:"handOffTimes,omitempty" yaml:"handOffTimes,omitempty"`
}
