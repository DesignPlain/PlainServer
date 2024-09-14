package types

type Ssm_getContactsRotationRecurrence struct {
	//
	NumberOfOnCalls int `json:"numberOfOnCalls,omitempty" yaml:"numberOfOnCalls,omitempty"`

	//
	RecurrenceMultiplier int `json:"recurrenceMultiplier,omitempty" yaml:"recurrenceMultiplier,omitempty"`

	//
	ShiftCoverages []Ssm_getContactsRotationRecurrenceShiftCoverage `json:"shiftCoverages,omitempty" yaml:"shiftCoverages,omitempty"`

	//
	WeeklySettings []Ssm_getContactsRotationRecurrenceWeeklySetting `json:"weeklySettings,omitempty" yaml:"weeklySettings,omitempty"`

	//
	DailySettings []Ssm_getContactsRotationRecurrenceDailySetting `json:"dailySettings,omitempty" yaml:"dailySettings,omitempty"`

	//
	MonthlySettings []Ssm_getContactsRotationRecurrenceMonthlySetting `json:"monthlySettings,omitempty" yaml:"monthlySettings,omitempty"`
}
