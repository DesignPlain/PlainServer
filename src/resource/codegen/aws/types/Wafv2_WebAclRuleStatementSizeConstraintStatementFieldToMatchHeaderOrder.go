package types

type Wafv2_WebAclRuleStatementSizeConstraintStatementFieldToMatchHeaderOrder struct {
	// Oversize handling tells AWS WAF what to do with a web request when the request component that the rule inspects is over the limits. Valid values include the following: `CONTINUE`, `MATCH`, `NO_MATCH`. See the AWS [documentation](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-statement-oversize-handling.html) for more information.
	OversizeHandling string `json:"oversizeHandling,omitempty" yaml:"oversizeHandling,omitempty"`
}
