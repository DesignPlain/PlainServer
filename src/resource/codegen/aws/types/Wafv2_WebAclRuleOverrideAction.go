package types

type Wafv2_WebAclRuleOverrideAction struct {
	// Override the rule action setting to count (i.e., only count matches). Configured as an empty block `{}`.
	Count Wafv2_WebAclRuleOverrideActionCount `json:"count,omitempty" yaml:"count,omitempty"`

	// Don't override the rule action setting. Configured as an empty block `{}`.
	None Wafv2_WebAclRuleOverrideActionNone `json:"none,omitempty" yaml:"none,omitempty"`
}
