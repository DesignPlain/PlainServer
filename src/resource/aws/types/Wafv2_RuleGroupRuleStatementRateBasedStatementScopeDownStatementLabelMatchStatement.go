package types

type Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementLabelMatchStatement struct {
	// The string to match against.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Specify whether you want to match using the label name or just the namespace. Valid values are `LABEL` or `NAMESPACE`.
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`
}
