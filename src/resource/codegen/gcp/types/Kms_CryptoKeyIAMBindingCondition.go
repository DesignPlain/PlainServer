package types

type Kms_CryptoKeyIAMBindingCondition struct {
	/*
	   An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.

	   > --Warning:-- The provider considers the `role` and condition contents (`title`+`description`+`expression`) as the
	   identifier for the binding. This means that if any part of the condition is changed out-of-band, the provider will
	   consider it to be an entirely different resource and will treat it as such.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`

	// A title for the expression, i.e. a short string describing its purpose.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
}
