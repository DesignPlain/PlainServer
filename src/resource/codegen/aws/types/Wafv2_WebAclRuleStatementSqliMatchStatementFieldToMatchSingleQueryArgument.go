package types

type Wafv2_WebAclRuleStatementSqliMatchStatementFieldToMatchSingleQueryArgument struct {
	// Name of the query header to inspect. This setting must be provided as lower case characters.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
