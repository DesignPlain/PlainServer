package types

type Compute_RegionUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch struct {
	/*
	   The name of the HTTP header to match. For matching against the HTTP request's
	   authority, use a headerMatch with the header name ":authority". For matching a
	   request's method, use the headerName ":method".
	*/
	HeaderName string `json:"headerName,omitempty" yaml:"headerName,omitempty"`

	/*
	   If set to false, the headerMatch is considered a match if the match criteria
	   above are met. If set to true, the headerMatch is considered a match if the
	   match criteria above are NOT met. Defaults to false.
	*/
	InvertMatch bool `json:"invertMatch,omitempty" yaml:"invertMatch,omitempty"`

	/*
	   The value of the header must start with the contents of prefixMatch. Only one of
	   exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch
	   must be set.
	*/
	PrefixMatch string `json:"prefixMatch,omitempty" yaml:"prefixMatch,omitempty"`

	/*
	   A header with the contents of headerName must exist. The match takes place
	   whether or not the request's header has a value or not. Only one of exactMatch,
	   prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set.
	*/
	PresentMatch bool `json:"presentMatch,omitempty" yaml:"presentMatch,omitempty"`

	/*
	   The header value must be an integer and its value must be in the range specified
	   in rangeMatch. If the header does not contain an integer, number or is empty,
	   the match fails. For example for a range [-5, 0]
	   - -3 will match
	   - 0 will not match
	   - 0.25 will not match
	   - -3someString will not match.
	   Only one of exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or
	   rangeMatch must be set.
	   Structure is documented below.
	*/
	RangeMatch Compute_RegionUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch `json:"rangeMatch,omitempty" yaml:"rangeMatch,omitempty"`

	/*
	   The value of the header must match the regular expression specified in
	   regexMatch. For regular expression grammar, please see:
	   en.cppreference.com/w/cpp/regex/ecmascript  For matching against a port
	   specified in the HTTP request, use a headerMatch with headerName set to PORT and
	   a regular expression that satisfies the RFC2616 Host header's port specifier.
	   Only one of exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or
	   rangeMatch must be set.
	*/
	RegexMatch string `json:"regexMatch,omitempty" yaml:"regexMatch,omitempty"`

	/*
	   The value of the header must end with the contents of suffixMatch. Only one of
	   exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch
	   must be set.
	*/
	SuffixMatch string `json:"suffixMatch,omitempty" yaml:"suffixMatch,omitempty"`

	/*
	   The value should exactly match contents of exactMatch. Only one of exactMatch,
	   prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set.
	*/
	ExactMatch string `json:"exactMatch,omitempty" yaml:"exactMatch,omitempty"`
}
