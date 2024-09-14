package types

type Securityhub_AutomationRuleCriteriaCriticality struct {
	//
	Lt float64 `json:"lt,omitempty" yaml:"lt,omitempty"`

	// The less-than-equal condition to be applied to a single field when querying for findings, provided as a String.
	Lte float64 `json:"lte,omitempty" yaml:"lte,omitempty"`

	// The equal-to condition to be applied to a single field when querying for findings, provided as a String.
	Eq float64 `json:"eq,omitempty" yaml:"eq,omitempty"`

	//
	Gt float64 `json:"gt,omitempty" yaml:"gt,omitempty"`

	// The greater-than-equal condition to be applied to a single field when querying for findings, provided as a String.
	Gte float64 `json:"gte,omitempty" yaml:"gte,omitempty"`
}
