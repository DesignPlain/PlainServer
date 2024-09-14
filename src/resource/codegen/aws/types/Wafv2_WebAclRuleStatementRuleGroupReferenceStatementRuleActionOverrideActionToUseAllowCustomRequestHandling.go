package types

type Wafv2_WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseAllowCustomRequestHandling struct {
	// The `insert_header` blocks used to define HTTP headers added to the request. See `insert_header` below for details.
	InsertHeaders []Wafv2_WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseAllowCustomRequestHandlingInsertHeader `json:"insertHeaders,omitempty" yaml:"insertHeaders,omitempty"`
}
