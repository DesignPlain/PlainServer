package types

type Wafv2_RuleGroupRuleStatementSqliMatchStatementFieldToMatchCookiesMatchPattern struct {
	//
	ExcludedCookies []string `json:"excludedCookies,omitempty" yaml:"excludedCookies,omitempty"`

	//
	IncludedCookies []string `json:"includedCookies,omitempty" yaml:"includedCookies,omitempty"`

	// An empty configuration block that is used for inspecting all headers.
	All Wafv2_RuleGroupRuleStatementSqliMatchStatementFieldToMatchCookiesMatchPatternAll `json:"all,omitempty" yaml:"all,omitempty"`
}
