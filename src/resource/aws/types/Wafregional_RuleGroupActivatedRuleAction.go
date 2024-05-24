package types

type Wafregional_RuleGroupActivatedRuleAction struct {
	// The rule type, either `REGULAR`, `RATE_BASED`, or `GROUP`. Defaults to `REGULAR`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
