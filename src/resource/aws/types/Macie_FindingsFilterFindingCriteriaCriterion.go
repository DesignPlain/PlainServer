package types

type Macie_FindingsFilterFindingCriteriaCriterion struct {
	// The value for the property doesn't match (doesn't equal) the specified value. If you specify multiple values, Amazon Macie uses OR logic to join the values.
	Neqs []string `json:"neqs,omitempty" yaml:"neqs,omitempty"`

	// The value for the property exclusively matches (equals an exact match for) all the specified values. If you specify multiple values, Amazon Macie uses AND logic to join the values.
	EqExactMatches []string `json:"eqExactMatches,omitempty" yaml:"eqExactMatches,omitempty"`

	// The value for the property matches (equals) the specified value. If you specify multiple values, Amazon Macie uses OR logic to join the values.
	Eqs []string `json:"eqs,omitempty" yaml:"eqs,omitempty"`

	// The name of the field to be evaluated.
	Field string `json:"field,omitempty" yaml:"field,omitempty"`

	// The value for the property is greater than the specified value.
	Gt string `json:"gt,omitempty" yaml:"gt,omitempty"`

	// The value for the property is greater than or equal to the specified value.
	Gte string `json:"gte,omitempty" yaml:"gte,omitempty"`

	// The value for the property is less than the specified value.
	Lt string `json:"lt,omitempty" yaml:"lt,omitempty"`

	// The value for the property is less than or equal to the specified value.
	Lte string `json:"lte,omitempty" yaml:"lte,omitempty"`
}
