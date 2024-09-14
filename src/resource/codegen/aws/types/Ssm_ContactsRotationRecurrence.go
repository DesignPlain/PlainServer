package types

type Ssm_ContactsRotationRecurrence struct {
	//
	DailySettings []Ssm_ContactsRotationRecurrenceDailySetting `json:"dailySettings,omitempty" yaml:"dailySettings,omitempty"`

	// (Optional) Information about on-call rotations that recur monthly. See Monthly Settings for more details.
	MonthlySettings []Ssm_ContactsRotationRecurrenceMonthlySetting `json:"monthlySettings,omitempty" yaml:"monthlySettings,omitempty"`

	// (Required) The number of contacts, or shift team members designated to be on call concurrently during a shift.
	NumberOfOnCalls int `json:"numberOfOnCalls,omitempty" yaml:"numberOfOnCalls,omitempty"`

	// (Required) The number of days, weeks, or months a single rotation lasts.
	RecurrenceMultiplier int `json:"recurrenceMultiplier,omitempty" yaml:"recurrenceMultiplier,omitempty"`

	// (Optional) Information about the days of the week that the on-call rotation coverage includes. See Shift Coverages for more details.
	ShiftCoverages []Ssm_ContactsRotationRecurrenceShiftCoverage `json:"shiftCoverages,omitempty" yaml:"shiftCoverages,omitempty"`

	// (Optional) Information about on-call rotations that recur weekly. See Weekly Settings for more details.
	WeeklySettings []Ssm_ContactsRotationRecurrenceWeeklySetting `json:"weeklySettings,omitempty" yaml:"weeklySettings,omitempty"`
}
