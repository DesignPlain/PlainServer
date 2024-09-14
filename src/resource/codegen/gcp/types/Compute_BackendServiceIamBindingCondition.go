package types

type Compute_BackendServiceIamBindingCondition struct {
	//
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`

	// A title for the expression, i.e. a short string describing its purpose.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
}
