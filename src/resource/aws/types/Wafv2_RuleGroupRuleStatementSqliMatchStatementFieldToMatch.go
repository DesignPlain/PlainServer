package types

type Wafv2_RuleGroupRuleStatementSqliMatchStatementFieldToMatch struct {
	// Inspect the cookies in the web request. See Cookies below for details.
	Cookies Wafv2_RuleGroupRuleStatementSqliMatchStatementFieldToMatchCookies `json:"cookies,omitempty" yaml:"cookies,omitempty"`

	// Inspect the request headers. See Header Order below for details.
	HeaderOrders []Wafv2_RuleGroupRuleStatementSqliMatchStatementFieldToMatchHeaderOrder `json:"headerOrders,omitempty" yaml:"headerOrders,omitempty"`

	// Inspect the request headers. See Headers below for details.
	Headers []Wafv2_RuleGroupRuleStatementSqliMatchStatementFieldToMatchHeader `json:"headers,omitempty" yaml:"headers,omitempty"`

	//
	Ja3Fingerprint Wafv2_RuleGroupRuleStatementSqliMatchStatementFieldToMatchJa3Fingerprint `json:"ja3Fingerprint,omitempty" yaml:"ja3Fingerprint,omitempty"`

	// Inspect the HTTP method. The method indicates the type of operation that the request is asking the origin to perform.
	Method Wafv2_RuleGroupRuleStatementSqliMatchStatementFieldToMatchMethod `json:"method,omitempty" yaml:"method,omitempty"`

	// Inspect a single header. See Single Header below for details.
	SingleHeader Wafv2_RuleGroupRuleStatementSqliMatchStatementFieldToMatchSingleHeader `json:"singleHeader,omitempty" yaml:"singleHeader,omitempty"`

	// Inspect a single query argument. See Single Query Argument below for details.
	SingleQueryArgument Wafv2_RuleGroupRuleStatementSqliMatchStatementFieldToMatchSingleQueryArgument `json:"singleQueryArgument,omitempty" yaml:"singleQueryArgument,omitempty"`

	// Inspect the request body, which immediately follows the request headers.
	Body Wafv2_RuleGroupRuleStatementSqliMatchStatementFieldToMatchBody `json:"body,omitempty" yaml:"body,omitempty"`

	// Inspect the request URI path. This is the part of a web request that identifies a resource, for example, `/images/daily-ad.jpg`.
	UriPath Wafv2_RuleGroupRuleStatementSqliMatchStatementFieldToMatchUriPath `json:"uriPath,omitempty" yaml:"uriPath,omitempty"`

	// Inspect the request body as JSON. See JSON Body for details.
	JsonBody Wafv2_RuleGroupRuleStatementSqliMatchStatementFieldToMatchJsonBody `json:"jsonBody,omitempty" yaml:"jsonBody,omitempty"`

	// Inspect the query string. This is the part of a URL that appears after a `?` character, if any.
	QueryString Wafv2_RuleGroupRuleStatementSqliMatchStatementFieldToMatchQueryString `json:"queryString,omitempty" yaml:"queryString,omitempty"`

	// Inspect all query arguments.
	AllQueryArguments Wafv2_RuleGroupRuleStatementSqliMatchStatementFieldToMatchAllQueryArguments `json:"allQueryArguments,omitempty" yaml:"allQueryArguments,omitempty"`
}
