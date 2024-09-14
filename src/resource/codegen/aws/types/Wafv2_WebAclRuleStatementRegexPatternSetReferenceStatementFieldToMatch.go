package types

type Wafv2_WebAclRuleStatementRegexPatternSetReferenceStatementFieldToMatch struct {
	// Inspect a string containing the list of the request's header names, ordered as they appear in the web request that AWS WAF receives for inspection. See `header_order` below for details.
	HeaderOrders []Wafv2_WebAclRuleStatementRegexPatternSetReferenceStatementFieldToMatchHeaderOrder `json:"headerOrders,omitempty" yaml:"headerOrders,omitempty"`

	// Inspect the JA3 fingerprint. See `ja3_fingerprint` below for details.
	Ja3Fingerprint Wafv2_WebAclRuleStatementRegexPatternSetReferenceStatementFieldToMatchJa3Fingerprint `json:"ja3Fingerprint,omitempty" yaml:"ja3Fingerprint,omitempty"`

	// Inspect the request body as JSON. See `json_body` for details.
	JsonBody Wafv2_WebAclRuleStatementRegexPatternSetReferenceStatementFieldToMatchJsonBody `json:"jsonBody,omitempty" yaml:"jsonBody,omitempty"`

	// Inspect the HTTP method. The method indicates the type of operation that the request is asking the origin to perform.
	Method Wafv2_WebAclRuleStatementRegexPatternSetReferenceStatementFieldToMatchMethod `json:"method,omitempty" yaml:"method,omitempty"`

	// Inspect the query string. This is the part of a URL that appears after a `?` character, if any.
	QueryString Wafv2_WebAclRuleStatementRegexPatternSetReferenceStatementFieldToMatchQueryString `json:"queryString,omitempty" yaml:"queryString,omitempty"`

	// Inspect a single query argument. See `single_query_argument` below for details.
	SingleQueryArgument Wafv2_WebAclRuleStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument `json:"singleQueryArgument,omitempty" yaml:"singleQueryArgument,omitempty"`

	// Inspect all query arguments.
	AllQueryArguments Wafv2_WebAclRuleStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments `json:"allQueryArguments,omitempty" yaml:"allQueryArguments,omitempty"`

	// Inspect the request body, which immediately follows the request headers. See `body` below for details.
	Body Wafv2_WebAclRuleStatementRegexPatternSetReferenceStatementFieldToMatchBody `json:"body,omitempty" yaml:"body,omitempty"`

	// Inspect the cookies in the web request. See `cookies` below for details.
	Cookies Wafv2_WebAclRuleStatementRegexPatternSetReferenceStatementFieldToMatchCookies `json:"cookies,omitempty" yaml:"cookies,omitempty"`

	// Inspect the request headers. See `headers` below for details.
	Headers []Wafv2_WebAclRuleStatementRegexPatternSetReferenceStatementFieldToMatchHeader `json:"headers,omitempty" yaml:"headers,omitempty"`

	// Inspect a single header. See `single_header` below for details.
	SingleHeader Wafv2_WebAclRuleStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader `json:"singleHeader,omitempty" yaml:"singleHeader,omitempty"`

	// Inspect the request URI path. This is the part of a web request that identifies a resource, for example, `/images/daily-ad.jpg`.
	UriPath Wafv2_WebAclRuleStatementRegexPatternSetReferenceStatementFieldToMatchUriPath `json:"uriPath,omitempty" yaml:"uriPath,omitempty"`
}
