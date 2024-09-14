package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchJsonBodyMatchPattern struct {
	// An empty configuration block that is used for inspecting all headers.
	All Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementFieldToMatchJsonBodyMatchPatternAll `json:"all,omitempty" yaml:"all,omitempty"`

	//
	IncludedPaths []string `json:"includedPaths,omitempty" yaml:"includedPaths,omitempty"`
}
