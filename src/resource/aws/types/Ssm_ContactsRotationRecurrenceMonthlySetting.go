package types

type Ssm_ContactsRotationRecurrenceMonthlySetting struct {
	// (Required) The day of the month when monthly recurring on-call rotations begin.
	DayOfMonth int `json:"dayOfMonth,omitempty" yaml:"dayOfMonth,omitempty"`

	// (Required) The hand off time. See Hand Off Time for more details.
	HandOffTime Ssm_ContactsRotationRecurrenceMonthlySettingHandOffTime `json:"handOffTime,omitempty" yaml:"handOffTime,omitempty"`
}
