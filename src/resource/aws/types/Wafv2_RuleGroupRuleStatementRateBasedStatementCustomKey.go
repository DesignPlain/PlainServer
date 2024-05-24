package types

type Wafv2_RuleGroupRuleStatementRateBasedStatementCustomKey struct {
	// (Optional) Use the first IP address in an HTTP header as an aggregate key. See `forwarded_ip` below for details.
	ForwardedIp Wafv2_RuleGroupRuleStatementRateBasedStatementCustomKeyForwardedIp `json:"forwardedIp,omitempty" yaml:"forwardedIp,omitempty"`

	// (Optional) Use the value of a header in the request as an aggregate key. See RateLimit `header` below for details.
	Header Wafv2_RuleGroupRuleStatementRateBasedStatementCustomKeyHeader `json:"header,omitempty" yaml:"header,omitempty"`

	// (Optional) Use the specified label namespace as an aggregate key. See RateLimit `label_namespace` below for details.
	LabelNamespace Wafv2_RuleGroupRuleStatementRateBasedStatementCustomKeyLabelNamespace `json:"labelNamespace,omitempty" yaml:"labelNamespace,omitempty"`

	// (Optional) Use the specified query argument as an aggregate key. See RateLimit `query_argument` below for details.
	QueryArgument Wafv2_RuleGroupRuleStatementRateBasedStatementCustomKeyQueryArgument `json:"queryArgument,omitempty" yaml:"queryArgument,omitempty"`

	// Inspect the query string. This is the part of a URL that appears after a `?` character, if any.
	QueryString Wafv2_RuleGroupRuleStatementRateBasedStatementCustomKeyQueryString `json:"queryString,omitempty" yaml:"queryString,omitempty"`

	// (Optional) Use the value of a cookie in the request as an aggregate key. See RateLimit `cookie` below for details.
	Cookie Wafv2_RuleGroupRuleStatementRateBasedStatementCustomKeyCookie `json:"cookie,omitempty" yaml:"cookie,omitempty"`

	// (Optional) Use the request's HTTP method as an aggregate key. See RateLimit `http_method` below for details.
	HttpMethod Wafv2_RuleGroupRuleStatementRateBasedStatementCustomKeyHttpMethod `json:"httpMethod,omitempty" yaml:"httpMethod,omitempty"`

	// (Optional) Use the request's originating IP address as an aggregate key. See `RateLimit ip` below for details.
	Ip Wafv2_RuleGroupRuleStatementRateBasedStatementCustomKeyIp `json:"ip,omitempty" yaml:"ip,omitempty"`

	// Inspect the request URI path. This is the part of a web request that identifies a resource, for example, `/images/daily-ad.jpg`.
	UriPath Wafv2_RuleGroupRuleStatementRateBasedStatementCustomKeyUriPath `json:"uriPath,omitempty" yaml:"uriPath,omitempty"`
}
