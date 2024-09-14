package types

type Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementFieldToMatch struct {
	// Inspect the cookies in the web request. See Cookies below for details.
	Cookies Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementFieldToMatchCookies `json:"cookies,omitempty" yaml:"cookies,omitempty"`

	// Inspect the request headers. See Header Order below for details.
	HeaderOrders []Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementFieldToMatchHeaderOrder `json:"headerOrders,omitempty" yaml:"headerOrders,omitempty"`

	// Inspect the request headers. See Headers below for details.
	Headers []Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementFieldToMatchHeader `json:"headers,omitempty" yaml:"headers,omitempty"`

	//
	Ja3Fingerprint Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementFieldToMatchJa3Fingerprint `json:"ja3Fingerprint,omitempty" yaml:"ja3Fingerprint,omitempty"`

	// Inspect the HTTP method. The method indicates the type of operation that the request is asking the origin to perform.
	Method Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementFieldToMatchMethod `json:"method,omitempty" yaml:"method,omitempty"`

	// Inspect the query string. This is the part of a URL that appears after a `?` character, if any.
	QueryString Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementFieldToMatchQueryString `json:"queryString,omitempty" yaml:"queryString,omitempty"`

	// Inspect all query arguments.
	AllQueryArguments Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementFieldToMatchAllQueryArguments `json:"allQueryArguments,omitempty" yaml:"allQueryArguments,omitempty"`

	// Inspect the request body, which immediately follows the request headers.
	Body Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementFieldToMatchBody `json:"body,omitempty" yaml:"body,omitempty"`

	// Inspect the request body as JSON. See JSON Body for details.
	JsonBody Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementFieldToMatchJsonBody `json:"jsonBody,omitempty" yaml:"jsonBody,omitempty"`

	// Inspect a single header. See Single Header below for details.
	SingleHeader Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementFieldToMatchSingleHeader `json:"singleHeader,omitempty" yaml:"singleHeader,omitempty"`

	// Inspect a single query argument. See Single Query Argument below for details.
	SingleQueryArgument Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementFieldToMatchSingleQueryArgument `json:"singleQueryArgument,omitempty" yaml:"singleQueryArgument,omitempty"`

	// Inspect the request URI path. This is the part of a web request that identifies a resource, for example, `/images/daily-ad.jpg`.
	UriPath Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementFieldToMatchUriPath `json:"uriPath,omitempty" yaml:"uriPath,omitempty"`
}
