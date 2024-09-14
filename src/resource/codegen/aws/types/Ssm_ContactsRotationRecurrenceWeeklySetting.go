package types

type Ssm_ContactsRotationRecurrenceWeeklySetting struct {
	// (Required) The day of the week when the shift coverage occurs.
	DayOfWeek string `json:"dayOfWeek,omitempty" yaml:"dayOfWeek,omitempty"`

	// (Required) The hand off time. See Hand Off Time for more details.
	HandOffTime Ssm_ContactsRotationRecurrenceWeeklySettingHandOffTime `json:"handOffTime,omitempty" yaml:"handOffTime,omitempty"`
}
