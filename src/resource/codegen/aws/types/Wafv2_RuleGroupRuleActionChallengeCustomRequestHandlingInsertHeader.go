package types

type Wafv2_RuleGroupRuleActionChallengeCustomRequestHandlingInsertHeader struct {
	// The value of the custom header.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// A friendly name of the rule group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
