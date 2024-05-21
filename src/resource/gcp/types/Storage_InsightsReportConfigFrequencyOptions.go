package types

type Storage_InsightsReportConfigFrequencyOptions struct {
	/*
	   The date to stop generating inventory reports. For example, {"day": 15, "month": 9, "year": 2022}.
	   Structure is documented below.
	*/
	EndDate Storage_InsightsReportConfigFrequencyOptionsEndDate `json:"endDate,omitempty" yaml:"endDate,omitempty"`

	/*
	   The frequency in which inventory reports are generated. Values are DAILY or WEEKLY.
	   Possible values are: `DAILY`, `WEEKLY`.
	*/
	Frequency string `json:"frequency,omitempty" yaml:"frequency,omitempty"`

	/*
	   The date to start generating inventory reports. For example, {"day": 15, "month": 8, "year": 2022}.
	   Structure is documented below.
	*/
	StartDate Storage_InsightsReportConfigFrequencyOptionsStartDate `json:"startDate,omitempty" yaml:"startDate,omitempty"`
}
