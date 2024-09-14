package types

type Wafv2_WebAclRuleStatementRateBasedStatement struct {
	// Limit on requests per 5-minute period for a single originating IP address.
	Limit int `json:"limit,omitempty" yaml:"limit,omitempty"`

	// Optional nested statement that narrows the scope of the rate-based statement to matching web requests. This can be any nestable statement, and you can nest statements at any level below this scope-down statement. See `statement` above for details. If `aggregate_key_type` is set to `CONSTANT`, this block is required.
	ScopeDownStatement Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatement `json:"scopeDownStatement,omitempty" yaml:"scopeDownStatement,omitempty"`

	// Setting that indicates how to aggregate the request counts. Valid values include: `CONSTANT`, `CUSTOM_KEYS`, `FORWARDED_IP`, or `IP`. Default: `IP`.
	AggregateKeyType string `json:"aggregateKeyType,omitempty" yaml:"aggregateKeyType,omitempty"`

	// Aggregate the request counts using one or more web request components as the aggregate keys. See `custom_key` below for details.
	CustomKeys []Wafv2_WebAclRuleStatementRateBasedStatementCustomKey `json:"customKeys,omitempty" yaml:"customKeys,omitempty"`

	/*
	   The amount of time, in seconds, that AWS WAF should include in its request counts, looking back from the current time. Valid values are `60`, `120`, `300`, and `600`. Defaults to `300` (5 minutes).

	   --NOTE:-- This setting doesn't determine how often AWS WAF checks the rate, but how far back it looks each time it checks. AWS WAF checks the rate about every 10 seconds.
	*/
	EvaluationWindowSec int `json:"evaluationWindowSec,omitempty" yaml:"evaluationWindowSec,omitempty"`

	// Configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin. If `aggregate_key_type` is set to `FORWARDED_IP`, this block is required. See `forwarded_ip_config` below for details.
	ForwardedIpConfig Wafv2_WebAclRuleStatementRateBasedStatementForwardedIpConfig `json:"forwardedIpConfig,omitempty" yaml:"forwardedIpConfig,omitempty"`
}
