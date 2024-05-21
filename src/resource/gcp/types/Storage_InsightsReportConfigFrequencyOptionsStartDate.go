package types

type Storage_InsightsReportConfigFrequencyOptionsStartDate struct {
	// The day of the month to start generating inventory reports.
	Day int `json:"day,omitempty" yaml:"day,omitempty"`

	// The month to start generating inventory reports.
	Month int `json:"month,omitempty" yaml:"month,omitempty"`

	// The year to start generating inventory reports
	Year int `json:"year,omitempty" yaml:"year,omitempty"`
}
