package types

type Waf_ByteMatchSetByteMatchTupleFieldToMatch struct {
	/*
	   When `type` is `HEADER`, enter the name of the header that you want to search, e.g., `User-Agent` or `Referer`.
	   If `type` is any other value, omit this field.
	*/
	Data string `json:"data,omitempty" yaml:"data,omitempty"`

	/*
	   The part of the web request that you want AWS WAF to search for a specified string.
	   e.g., `HEADER`, `METHOD` or `BODY`.
	   See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html)
	   for all supported values.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
