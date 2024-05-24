package types

type Customerprofiles_DomainMatchingJobSchedule struct {
	// The day when the Identity Resolution Job should run every week.
	DayOfTheWeek string `json:"dayOfTheWeek,omitempty" yaml:"dayOfTheWeek,omitempty"`

	// The time when the Identity Resolution Job should run every week.
	Time string `json:"time,omitempty" yaml:"time,omitempty"`
}
