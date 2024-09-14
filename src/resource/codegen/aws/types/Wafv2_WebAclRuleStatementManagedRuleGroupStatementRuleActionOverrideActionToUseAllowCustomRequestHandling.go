package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseAllowCustomRequestHandling struct {
	// The `insert_header` blocks used to define HTTP headers added to the request. See `insert_header` below for details.
	InsertHeaders []Wafv2_WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseAllowCustomRequestHandlingInsertHeader `json:"insertHeaders,omitempty" yaml:"insertHeaders,omitempty"`
}
