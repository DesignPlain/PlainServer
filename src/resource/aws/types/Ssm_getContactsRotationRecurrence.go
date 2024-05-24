package types

type Ssm_getContactsRotationRecurrence struct {
	//
	RecurrenceMultiplier int `json:"recurrenceMultiplier,omitempty" yaml:"recurrenceMultiplier,omitempty"`

	//
	ShiftCoverages []string `json:"shiftCoverages,omitempty" yaml:"shiftCoverages,omitempty"`

	//
	WeeklySettings []string `json:"weeklySettings,omitempty" yaml:"weeklySettings,omitempty"`

	//
	DailySettings []string `json:"dailySettings,omitempty" yaml:"dailySettings,omitempty"`

	//
	MonthlySettings []string `json:"monthlySettings,omitempty" yaml:"monthlySettings,omitempty"`

	//
	NumberOfOnCalls int `json:"numberOfOnCalls,omitempty" yaml:"numberOfOnCalls,omitempty"`
}
