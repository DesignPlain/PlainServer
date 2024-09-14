package types

type Wafregional_ByteMatchSetByteMatchTupleFieldToMatch struct {
	// When the value of Type is HEADER, enter the name of the header that you want AWS WAF to search, for example, User-Agent or Referer. If the value of Type is any other value, omit Data.
	Data string `json:"data,omitempty" yaml:"data,omitempty"`

	// The part of the web request that you want AWS WAF to search for a specified string.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
