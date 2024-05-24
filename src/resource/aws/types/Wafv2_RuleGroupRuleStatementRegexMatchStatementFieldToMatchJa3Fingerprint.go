package types

type Wafv2_RuleGroupRuleStatementRegexMatchStatementFieldToMatchJa3Fingerprint struct {
	// The match status to assign to the web request if the request doesn't have a valid IP address in the specified position. Valid values include: `MATCH` or `NO_MATCH`.
	FallbackBehavior string `json:"fallbackBehavior,omitempty" yaml:"fallbackBehavior,omitempty"`
}
