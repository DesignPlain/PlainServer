package types

type Wafv2_WebAclRuleStatementByteMatchStatementFieldToMatchCookiesMatchPattern struct {
	//
	ExcludedCookies []string `json:"excludedCookies,omitempty" yaml:"excludedCookies,omitempty"`

	//
	IncludedCookies []string `json:"includedCookies,omitempty" yaml:"includedCookies,omitempty"`

	// An empty configuration block that is used for inspecting all headers.
	All Wafv2_WebAclRuleStatementByteMatchStatementFieldToMatchCookiesMatchPatternAll `json:"all,omitempty" yaml:"all,omitempty"`
}
