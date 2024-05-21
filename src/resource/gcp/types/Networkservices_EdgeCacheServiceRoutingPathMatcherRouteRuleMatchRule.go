package types

type Networkservices_EdgeCacheServiceRoutingPathMatcherRouteRuleMatchRule struct {
	/*
	   For satisfying the matchRule condition, the path of the request
	   must match the wildcard pattern specified in pathTemplateMatch
	   after removing any query parameters and anchor that may be part
	   of the original URL.
	   pathTemplateMatch must be between 1 and 255 characters
	   (inclusive).  The pattern specified by pathTemplateMatch may
	   have at most 5 wildcard operators and at most 5 variable
	   captures in total.
	*/
	PathTemplateMatch string `json:"pathTemplateMatch,omitempty" yaml:"pathTemplateMatch,omitempty"`

	// For satisfying the matchRule condition, the request's path must begin with the specified prefixMatch. prefixMatch must begin with a /.
	PrefixMatch string `json:"prefixMatch,omitempty" yaml:"prefixMatch,omitempty"`

	/*
	   Specifies a list of query parameter match criteria, all of which must match corresponding query parameters in the request.
	   Structure is documented below.
	*/
	QueryParameterMatches []Networkservices_EdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleQueryParameterMatch `json:"queryParameterMatches,omitempty" yaml:"queryParameterMatches,omitempty"`

	// For satisfying the matchRule condition, the path of the request must exactly match the value specified in fullPathMatch after removing any query parameters and anchor that may be part of the original URL.
	FullPathMatch string `json:"fullPathMatch,omitempty" yaml:"fullPathMatch,omitempty"`

	/*
	   Specifies a list of header match criteria, all of which must match corresponding headers in the request.
	   Structure is documented below.
	*/
	HeaderMatches []Networkservices_EdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleHeaderMatch `json:"headerMatches,omitempty" yaml:"headerMatches,omitempty"`

	// Specifies that prefixMatch and fullPathMatch matches are case sensitive.
	IgnoreCase bool `json:"ignoreCase,omitempty" yaml:"ignoreCase,omitempty"`
}
