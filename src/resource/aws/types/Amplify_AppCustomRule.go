package types

type Amplify_AppCustomRule struct {
	// Condition for a URL rewrite or redirect rule, such as a country code.
	Condition string `json:"condition,omitempty" yaml:"condition,omitempty"`

	// Source pattern for a URL rewrite or redirect rule.
	Source string `json:"source,omitempty" yaml:"source,omitempty"`

	// Status code for a URL rewrite or redirect rule. Valid values: `200`, `301`, `302`, `404`, `404-200`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Target pattern for a URL rewrite or redirect rule.
	Target string `json:"target,omitempty" yaml:"target,omitempty"`
}
