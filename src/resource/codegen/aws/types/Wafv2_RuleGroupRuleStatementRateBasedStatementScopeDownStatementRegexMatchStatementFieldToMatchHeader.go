package types

type Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchHeader struct {
	// The filter to use to identify the subset of headers to inspect in a web request. The `match_pattern` block supports only one of the following arguments:
	MatchPattern Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchHeaderMatchPattern `json:"matchPattern,omitempty" yaml:"matchPattern,omitempty"`

	// The parts of the headers to inspect with the rule inspection criteria. If you specify `All`, AWS WAF inspects both keys and values. Valid values include the following: `ALL`, `Key`, `Value`.
	MatchScope string `json:"matchScope,omitempty" yaml:"matchScope,omitempty"`

	// Oversize handling tells AWS WAF what to do with a web request when the request component that the rule inspects is over the limits. Valid values include the following: `CONTINUE`, `MATCH`, `NO_MATCH`. See the AWS [documentation](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-statement-oversize-handling.html) for more information.
	OversizeHandling string `json:"oversizeHandling,omitempty" yaml:"oversizeHandling,omitempty"`
}
