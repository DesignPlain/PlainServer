package types

type Storage_InsightsReportConfigFrequencyOptionsEndDate struct {
	// The day of the month to stop generating inventory reports.
	Day int `json:"day,omitempty" yaml:"day,omitempty"`

	// The month to stop generating inventory reports.
	Month int `json:"month,omitempty" yaml:"month,omitempty"`

	// The year to stop generating inventory reports
	Year int `json:"year,omitempty" yaml:"year,omitempty"`
}
