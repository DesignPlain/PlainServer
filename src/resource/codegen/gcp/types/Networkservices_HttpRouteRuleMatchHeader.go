package types

type Networkservices_HttpRouteRuleMatchHeader struct {
	// If specified, the match result will be inverted before checking. Default value is set to false.
	InvertMatch bool `json:"invertMatch,omitempty" yaml:"invertMatch,omitempty"`

	// The value of the header must start with the contents of prefixMatch.
	PrefixMatch string `json:"prefixMatch,omitempty" yaml:"prefixMatch,omitempty"`

	// A header with headerName must exist. The match takes place whether or not the header has a value.
	PresentMatch bool `json:"presentMatch,omitempty" yaml:"presentMatch,omitempty"`

	/*
	   If specified, the rule will match if the request header value is within the range.
	   Structure is documented below.
	*/
	RangeMatch Networkservices_HttpRouteRuleMatchHeaderRangeMatch `json:"rangeMatch,omitempty" yaml:"rangeMatch,omitempty"`

	// The value of the header must match the regular expression specified in regexMatch.
	RegexMatch string `json:"regexMatch,omitempty" yaml:"regexMatch,omitempty"`

	// The value of the header must end with the contents of suffixMatch.
	SuffixMatch string `json:"suffixMatch,omitempty" yaml:"suffixMatch,omitempty"`

	// The value of the header should match exactly the content of exactMatch.
	ExactMatch string `json:"exactMatch,omitempty" yaml:"exactMatch,omitempty"`

	// The name of the HTTP header to match against.
	Header string `json:"header,omitempty" yaml:"header,omitempty"`
}
