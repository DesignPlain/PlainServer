package types

type Waf_SqlInjectionMatchSetSqlInjectionMatchTuple struct {
	// Specifies where in a web request to look for snippets of malicious SQL code.
	FieldToMatch Waf_SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch `json:"fieldToMatch,omitempty" yaml:"fieldToMatch,omitempty"`

	/*
	   Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	   If you specify a transformation, AWS WAF performs the transformation on `field_to_match` before inspecting a request for a match.
	   e.g., `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
	   See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_SqlInjectionMatchTuple.html#WAF-Type-SqlInjectionMatchTuple-TextTransformation)
	   for all supported values.
	*/
	TextTransformation string `json:"textTransformation,omitempty" yaml:"textTransformation,omitempty"`
}
