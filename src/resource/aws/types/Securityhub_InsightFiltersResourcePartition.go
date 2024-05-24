package types

type Securityhub_InsightFiltersResourcePartition struct {
	// The condition to apply to a string value when querying for findings. Valid values include: `EQUALS` and `NOT_EQUALS`.
	Comparison string `json:"comparison,omitempty" yaml:"comparison,omitempty"`

	// A date range value for the date filter, provided as an Integer.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
