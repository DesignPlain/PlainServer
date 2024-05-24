package types

type Accessanalyzer_ArchiveRuleFilter struct {
	// Filter criteria.
	Criteria string `json:"criteria,omitempty" yaml:"criteria,omitempty"`

	// Equals comparator.
	Eqs []string `json:"eqs,omitempty" yaml:"eqs,omitempty"`

	// Boolean comparator.
	Exists string `json:"exists,omitempty" yaml:"exists,omitempty"`

	// Not Equals comparator.
	Neqs []string `json:"neqs,omitempty" yaml:"neqs,omitempty"`

	// Contains comparator.
	Contains []string `json:"contains,omitempty" yaml:"contains,omitempty"`
}
