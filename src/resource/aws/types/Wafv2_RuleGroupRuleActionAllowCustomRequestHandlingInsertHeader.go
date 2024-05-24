package types

type Wafv2_RuleGroupRuleActionAllowCustomRequestHandlingInsertHeader struct {
	// A friendly name of the rule group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The value of the custom header.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
