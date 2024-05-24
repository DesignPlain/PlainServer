package types

type Wafregional_ByteMatchSetByteMatchTuple struct {
	// Settings for the ByteMatchTuple. FieldToMatch documented below.
	FieldToMatch Wafregional_ByteMatchSetByteMatchTupleFieldToMatch `json:"fieldToMatch,omitempty" yaml:"fieldToMatch,omitempty"`

	// Within the portion of a web request that you want to search.
	PositionalConstraint string `json:"positionalConstraint,omitempty" yaml:"positionalConstraint,omitempty"`

	// The value that you want AWS WAF to search for. The maximum length of the value is 50 bytes.
	TargetString string `json:"targetString,omitempty" yaml:"targetString,omitempty"`

	/*
	   The formatting way for web request.

	   FieldToMatch(field_to_match) support following:
	*/
	TextTransformation string `json:"textTransformation,omitempty" yaml:"textTransformation,omitempty"`
}
