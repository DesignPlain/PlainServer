package types

type Guardduty_FilterFindingCriteriaCriterion struct {
	// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	LessThanOrEqual string `json:"lessThanOrEqual,omitempty" yaml:"lessThanOrEqual,omitempty"`

	// List of string values to be evaluated.
	NotEquals []string `json:"notEquals,omitempty" yaml:"notEquals,omitempty"`

	// List of string values to be evaluated.
	Equals []string `json:"equals,omitempty" yaml:"equals,omitempty"`

	// The name of the field to be evaluated. The full list of field names can be found in [AWS documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_filter-findings.html#filter_criteria).
	Field string `json:"field,omitempty" yaml:"field,omitempty"`

	// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	GreaterThan string `json:"greaterThan,omitempty" yaml:"greaterThan,omitempty"`

	// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	GreaterThanOrEqual string `json:"greaterThanOrEqual,omitempty" yaml:"greaterThanOrEqual,omitempty"`

	// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	LessThan string `json:"lessThan,omitempty" yaml:"lessThan,omitempty"`
}
