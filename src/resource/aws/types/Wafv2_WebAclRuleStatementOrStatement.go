package types

type Wafv2_WebAclRuleStatementOrStatement struct {
	// The statements to combine.
	Statements []Wafv2_WebAclRuleStatement `json:"statements,omitempty" yaml:"statements,omitempty"`
}
