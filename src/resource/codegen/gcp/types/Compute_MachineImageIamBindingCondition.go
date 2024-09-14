package types

type Compute_MachineImageIamBindingCondition struct {
	// A title for the expression, i.e. a short string describing its purpose.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`

	// An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`
}
