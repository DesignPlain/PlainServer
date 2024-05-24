package types

type Wafv2_WebAclRuleStatementRateBasedStatementCustomKey struct {
	// Use the request's URI path as an aggregate key. See RateLimit `uri_path` below for details.
	UriPath Wafv2_WebAclRuleStatementRateBasedStatementCustomKeyUriPath `json:"uriPath,omitempty" yaml:"uriPath,omitempty"`

	// Use the first IP address in an HTTP header as an aggregate key. See `forwarded_ip` below for details.
	ForwardedIp Wafv2_WebAclRuleStatementRateBasedStatementCustomKeyForwardedIp `json:"forwardedIp,omitempty" yaml:"forwardedIp,omitempty"`

	// Use the request's HTTP method as an aggregate key. See RateLimit `http_method` below for details.
	HttpMethod Wafv2_WebAclRuleStatementRateBasedStatementCustomKeyHttpMethod `json:"httpMethod,omitempty" yaml:"httpMethod,omitempty"`

	// Use the request's originating IP address as an aggregate key. See `RateLimit ip` below for details.
	Ip Wafv2_WebAclRuleStatementRateBasedStatementCustomKeyIp `json:"ip,omitempty" yaml:"ip,omitempty"`

	// Use the specified label namespace as an aggregate key. See RateLimit `label_namespace` below for details.
	LabelNamespace Wafv2_WebAclRuleStatementRateBasedStatementCustomKeyLabelNamespace `json:"labelNamespace,omitempty" yaml:"labelNamespace,omitempty"`

	// Use the specified query argument as an aggregate key. See RateLimit `query_argument` below for details.
	QueryArgument Wafv2_WebAclRuleStatementRateBasedStatementCustomKeyQueryArgument `json:"queryArgument,omitempty" yaml:"queryArgument,omitempty"`

	// Use the request's query string as an aggregate key. See RateLimit `query_string` below for details.
	QueryString Wafv2_WebAclRuleStatementRateBasedStatementCustomKeyQueryString `json:"queryString,omitempty" yaml:"queryString,omitempty"`

	// Use the value of a cookie in the request as an aggregate key. See RateLimit `cookie` below for details.
	Cookie Wafv2_WebAclRuleStatementRateBasedStatementCustomKeyCookie `json:"cookie,omitempty" yaml:"cookie,omitempty"`

	// Use the value of a header in the request as an aggregate key. See RateLimit `header` below for details.
	Header Wafv2_WebAclRuleStatementRateBasedStatementCustomKeyHeader `json:"header,omitempty" yaml:"header,omitempty"`
}
