package types

type Wafregional_XssMatchSetXssMatchTuple struct {
	// Specifies where in a web request to look for cross-site scripting attacks.
	FieldToMatch Wafregional_XssMatchSetXssMatchTupleFieldToMatch `json:"fieldToMatch,omitempty" yaml:"fieldToMatch,omitempty"`

	// Which text transformation, if any, to perform on the web request before inspecting the request for cross-site scripting attacks.
	TextTransformation string `json:"textTransformation,omitempty" yaml:"textTransformation,omitempty"`
}
