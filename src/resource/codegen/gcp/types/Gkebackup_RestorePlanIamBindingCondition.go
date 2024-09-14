package types

type Gkebackup_RestorePlanIamBindingCondition struct {
	// User specified descriptive string for this RestorePlan.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	//
	Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`

	//
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
}
