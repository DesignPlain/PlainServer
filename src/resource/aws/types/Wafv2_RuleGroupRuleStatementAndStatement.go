package types

type Wafv2_RuleGroupRuleStatementAndStatement struct {
	// The statements to combine.
	Statements []Wafv2_RuleGroupRuleStatement `json:"statements,omitempty" yaml:"statements,omitempty"`
}
