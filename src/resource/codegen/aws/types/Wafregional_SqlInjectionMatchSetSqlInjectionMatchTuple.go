package types

type Wafregional_SqlInjectionMatchSetSqlInjectionMatchTuple struct {
	// Specifies where in a web request to look for snippets of malicious SQL code.
	FieldToMatch Wafregional_SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch `json:"fieldToMatch,omitempty" yaml:"fieldToMatch,omitempty"`

	/*
	   Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	   If you specify a transformation, AWS WAF performs the transformation on `field_to_match` before inspecting a request for a match.
	   e.g., `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
	   See [docs](https://docs.aws.amazon.com/waf/latest/APIReference/API_regional_SqlInjectionMatchTuple.html#WAF-Type-regional_SqlInjectionMatchTuple-TextTransformation)
	   for all supported values.
	*/
	TextTransformation string `json:"textTransformation,omitempty" yaml:"textTransformation,omitempty"`
}
