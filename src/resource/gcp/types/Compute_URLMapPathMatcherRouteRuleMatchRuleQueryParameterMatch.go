package types

type Compute_URLMapPathMatcherRouteRuleMatchRuleQueryParameterMatch struct {
	/*
	   The queryParameterMatch matches if the value of the parameter exactly matches
	   the contents of exactMatch. Only one of presentMatch, exactMatch and regexMatch
	   must be set.
	*/
	ExactMatch string `json:"exactMatch,omitempty" yaml:"exactMatch,omitempty"`

	/*
	   The name of the query parameter to match. The query parameter must exist in the
	   request, in the absence of which the request match fails.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Specifies that the queryParameterMatch matches if the request contains the query
	   parameter, irrespective of whether the parameter has a value or not. Only one of
	   presentMatch, exactMatch and regexMatch must be set.
	*/
	PresentMatch bool `json:"presentMatch,omitempty" yaml:"presentMatch,omitempty"`

	/*
	   The queryParameterMatch matches if the value of the parameter matches the
	   regular expression specified by regexMatch. For the regular expression grammar,
	   please see en.cppreference.com/w/cpp/regex/ecmascript  Only one of presentMatch,
	   exactMatch and regexMatch must be set.
	*/
	RegexMatch string `json:"regexMatch,omitempty" yaml:"regexMatch,omitempty"`
}
