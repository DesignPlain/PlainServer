package types

type Lb_ListenerRuleCondition struct {
	// Contains a single `values` item which is a list of host header patterns to match. The maximum size of each pattern is 128 characters. Comparison is case insensitive. Wildcard characters supported: - (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied.
	HostHeader Lb_ListenerRuleConditionHostHeader `json:"hostHeader,omitempty" yaml:"hostHeader,omitempty"`

	// HTTP headers to match. HTTP Header block fields documented below.
	HttpHeader Lb_ListenerRuleConditionHttpHeader `json:"httpHeader,omitempty" yaml:"httpHeader,omitempty"`

	// Contains a single `values` item which is a list of HTTP request methods or verbs to match. Maximum size is 40 characters. Only allowed characters are A-Z, hyphen (-) and underscore (\_). Comparison is case sensitive. Wildcards are not supported. Only one needs to match for the condition to be satisfied. AWS recommends that GET and HEAD requests are routed in the same way because the response to a HEAD request may be cached.
	HttpRequestMethod Lb_ListenerRuleConditionHttpRequestMethod `json:"httpRequestMethod,omitempty" yaml:"httpRequestMethod,omitempty"`

	// Contains a single `values` item which is a list of path patterns to match against the request URL. Maximum size of each pattern is 128 characters. Comparison is case sensitive. Wildcard characters supported: - (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied. Path pattern is compared only to the path of the URL, not to its query string. To compare against the query string, use a `query_string` condition.
	PathPattern Lb_ListenerRuleConditionPathPattern `json:"pathPattern,omitempty" yaml:"pathPattern,omitempty"`

	// Query strings to match. Query String block fields documented below.
	QueryStrings []Lb_ListenerRuleConditionQueryString `json:"queryStrings,omitempty" yaml:"queryStrings,omitempty"`

	/*
	   Contains a single `values` item which is a list of source IP CIDR notations to match. You can use both IPv4 and IPv6 addresses. Wildcards are not supported. Condition is satisfied if the source IP address of the request matches one of the CIDR blocks. Condition is not satisfied by the addresses in the `X-Forwarded-For` header, use `http_header` condition instead.

	   > --NOTE::-- Exactly one of `host_header`, `http_header`, `http_request_method`, `path_pattern`, `query_string` or `source_ip` must be set per condition.
	*/
	SourceIp Lb_ListenerRuleConditionSourceIp `json:"sourceIp,omitempty" yaml:"sourceIp,omitempty"`
}
