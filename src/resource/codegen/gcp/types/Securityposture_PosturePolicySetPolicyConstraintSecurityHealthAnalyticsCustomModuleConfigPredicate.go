package types

type Securityposture_PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModuleConfigPredicate struct {
	// Title for the expression, i.e. a short string describing its purpose.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`

	// Description of the expression
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`

	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
