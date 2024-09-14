package types

type Wafv2_RuleGroupRuleActionCountCustomRequestHandling struct {
	// The `insert_header` blocks used to define HTTP headers added to the request. See Custom HTTP Header below for details.
	InsertHeaders []Wafv2_RuleGroupRuleActionCountCustomRequestHandlingInsertHeader `json:"insertHeaders,omitempty" yaml:"insertHeaders,omitempty"`
}
