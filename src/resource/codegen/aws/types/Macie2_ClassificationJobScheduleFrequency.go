package types

type Macie2_ClassificationJobScheduleFrequency struct {
	// Specifies a weekly recurrence pattern for running the job.
	WeeklySchedule string `json:"weeklySchedule,omitempty" yaml:"weeklySchedule,omitempty"`

	// Specifies a daily recurrence pattern for running the job.
	DailySchedule bool `json:"dailySchedule,omitempty" yaml:"dailySchedule,omitempty"`

	// Specifies a monthly recurrence pattern for running the job.
	MonthlySchedule int `json:"monthlySchedule,omitempty" yaml:"monthlySchedule,omitempty"`
}
