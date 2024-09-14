package types

type Macie2_ClassificationJobS3JobDefinitionScopingIncludesAndSimpleScopeTerm struct {
	// The object property to use in the condition.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// An array that lists the values to use in the condition.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// The operator to use in a condition. Valid values are: `EQ`, `GT`, `GTE`, `LT`, `LTE`, `NE`, `CONTAINS`, `STARTS_WITH`
	Comparator string `json:"comparator,omitempty" yaml:"comparator,omitempty"`
}
