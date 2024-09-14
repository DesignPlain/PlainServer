package types

type Securityhub_InsightFiltersNetworkDestinationPort struct {
	// The equal-to condition to be applied to a single field when querying for findings, provided as a String.
	Eq string `json:"eq,omitempty" yaml:"eq,omitempty"`

	// The greater-than-equal condition to be applied to a single field when querying for findings, provided as a String.
	Gte string `json:"gte,omitempty" yaml:"gte,omitempty"`

	// The less-than-equal condition to be applied to a single field when querying for findings, provided as a String.
	Lte string `json:"lte,omitempty" yaml:"lte,omitempty"`
}
