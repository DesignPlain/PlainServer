package types

type Wafv2_RuleGroupRuleStatementOrStatement struct {
	// The statements to combine.
	Statements []Wafv2_RuleGroupRuleStatement `json:"statements,omitempty" yaml:"statements,omitempty"`
}
