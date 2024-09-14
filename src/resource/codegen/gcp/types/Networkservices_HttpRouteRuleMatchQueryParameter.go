package types

type Networkservices_HttpRouteRuleMatchQueryParameter struct {
	// The value of the query parameter must exactly match the contents of exactMatch.
	ExactMatch string `json:"exactMatch,omitempty" yaml:"exactMatch,omitempty"`

	// Specifies that the QueryParameterMatcher matches if request contains query parameter, irrespective of whether the parameter has a value or not.
	PresentMatch bool `json:"presentMatch,omitempty" yaml:"presentMatch,omitempty"`

	// The name of the query parameter to match.
	QueryParameter string `json:"queryParameter,omitempty" yaml:"queryParameter,omitempty"`

	// The value of the query parameter must match the regular expression specified by regexMatch.For regular expression grammar, please see https://github.com/google/re2/wiki/Syntax
	RegexMatch string `json:"regexMatch,omitempty" yaml:"regexMatch,omitempty"`
}
