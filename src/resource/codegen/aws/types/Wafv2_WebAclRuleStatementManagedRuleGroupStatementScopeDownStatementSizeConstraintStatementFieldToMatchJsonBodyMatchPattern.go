package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementSizeConstraintStatementFieldToMatchJsonBodyMatchPattern struct {
	//
	IncludedPaths []string `json:"includedPaths,omitempty" yaml:"includedPaths,omitempty"`

	// An empty configuration block that is used for inspecting all headers.
	All Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementSizeConstraintStatementFieldToMatchJsonBodyMatchPatternAll `json:"all,omitempty" yaml:"all,omitempty"`
}
