package types

type Networkservices_HttpRouteRuleMatch struct {
	// The HTTP request path value should exactly match this value.
	FullPathMatch string `json:"fullPathMatch,omitempty" yaml:"fullPathMatch,omitempty"`

	/*
	   Specifies a list of HTTP request headers to match against.
	   Structure is documented below.
	*/
	Headers []Networkservices_HttpRouteRuleMatchHeader `json:"headers,omitempty" yaml:"headers,omitempty"`

	// Specifies if prefixMatch and fullPathMatch matches are case sensitive. The default value is false.
	IgnoreCase bool `json:"ignoreCase,omitempty" yaml:"ignoreCase,omitempty"`

	// The HTTP request path value must begin with specified prefixMatch. prefixMatch must begin with a /.
	PrefixMatch string `json:"prefixMatch,omitempty" yaml:"prefixMatch,omitempty"`

	/*
	   Specifies a list of query parameters to match against.
	   Structure is documented below.
	*/
	QueryParameters []Networkservices_HttpRouteRuleMatchQueryParameter `json:"queryParameters,omitempty" yaml:"queryParameters,omitempty"`

	// The HTTP request path value must satisfy the regular expression specified by regexMatch after removing any query parameters and anchor supplied with the original URL. For regular expression grammar, please see https://github.com/google/re2/wiki/Syntax
	RegexMatch string `json:"regexMatch,omitempty" yaml:"regexMatch,omitempty"`
}
