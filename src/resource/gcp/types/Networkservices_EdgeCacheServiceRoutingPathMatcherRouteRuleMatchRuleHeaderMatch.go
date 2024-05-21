package types

type Networkservices_EdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleHeaderMatch struct {
	// The header name to match on.
	HeaderName string `json:"headerName,omitempty" yaml:"headerName,omitempty"`

	/*
	   If set to false (default), the headerMatch is considered a match if the match criteria above are met.
	   If set to true, the headerMatch is considered a match if the match criteria above are NOT met.
	*/
	InvertMatch bool `json:"invertMatch,omitempty" yaml:"invertMatch,omitempty"`

	// The value of the header must start with the contents of prefixMatch.
	PrefixMatch string `json:"prefixMatch,omitempty" yaml:"prefixMatch,omitempty"`

	// A header with the contents of headerName must exist. The match takes place whether or not the request's header has a value.
	PresentMatch bool `json:"presentMatch,omitempty" yaml:"presentMatch,omitempty"`

	// The value of the header must end with the contents of suffixMatch.
	SuffixMatch string `json:"suffixMatch,omitempty" yaml:"suffixMatch,omitempty"`

	// The value of the header should exactly match contents of exactMatch.
	ExactMatch string `json:"exactMatch,omitempty" yaml:"exactMatch,omitempty"`
}
