package types

type Wafregional_WebAclRuleAction struct {
	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. Valid values for `action` are `ALLOW`, `BLOCK` or `COUNT`. Valid values for `override_action` are `COUNT` and `NONE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
