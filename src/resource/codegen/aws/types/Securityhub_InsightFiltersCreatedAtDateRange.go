package types

type Securityhub_InsightFiltersCreatedAtDateRange struct {
	// A date range unit for the date filter. Valid values: `DAYS`.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	// A date range value for the date filter, provided as an Integer.
	Value int `json:"value,omitempty" yaml:"value,omitempty"`
}
