package types

type Wafv2_WebAclRuleStatementAndStatement struct {
	// The statements to combine.
	Statements []Wafv2_WebAclRuleStatement `json:"statements,omitempty" yaml:"statements,omitempty"`
}
