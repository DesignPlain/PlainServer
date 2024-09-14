package types

type Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatementFieldToMatchSingleHeader struct {
	// The name of the query header to inspect. This setting must be provided as lower case characters.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
