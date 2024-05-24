package types

type Wafv2_RuleGroupRuleStatementRateBasedStatement struct {
	// An optional nested statement that narrows the scope of the rate-based statement to matching web requests. This can be any nestable statement, and you can nest statements at any level below this scope-down statement. See Statement above for details. If `aggregate_key_type` is set to `CONSTANT`, this block is required.
	ScopeDownStatement Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatement `json:"scopeDownStatement,omitempty" yaml:"scopeDownStatement,omitempty"`

	// Setting that indicates how to aggregate the request counts. Valid values include: `CONSTANT`, `CUSTOM_KEYS`, `FORWARDED_IP` or `IP`. Default: `IP`.
	AggregateKeyType string `json:"aggregateKeyType,omitempty" yaml:"aggregateKeyType,omitempty"`

	// Aggregate the request counts using one or more web request components as the aggregate keys. See `custom_key` below for details.
	CustomKeys []Wafv2_RuleGroupRuleStatementRateBasedStatementCustomKey `json:"customKeys,omitempty" yaml:"customKeys,omitempty"`

	// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin. If `aggregate_key_type` is set to `FORWARDED_IP`, this block is required. See Forwarded IP Config below for details.
	ForwardedIpConfig Wafv2_RuleGroupRuleStatementRateBasedStatementForwardedIpConfig `json:"forwardedIpConfig,omitempty" yaml:"forwardedIpConfig,omitempty"`

	// The limit on requests per 5-minute period for a single originating IP address.
	Limit int `json:"limit,omitempty" yaml:"limit,omitempty"`
}
