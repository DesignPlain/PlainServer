package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchCookiesMatchPattern struct {
	// An empty configuration block that is used for inspecting all headers.
	All Wafv2_WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementFieldToMatchCookiesMatchPatternAll `json:"all,omitempty" yaml:"all,omitempty"`

	//
	ExcludedCookies []string `json:"excludedCookies,omitempty" yaml:"excludedCookies,omitempty"`

	//
	IncludedCookies []string `json:"includedCookies,omitempty" yaml:"includedCookies,omitempty"`
}
