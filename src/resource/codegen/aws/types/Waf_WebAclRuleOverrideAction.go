package types

type Waf_WebAclRuleOverrideAction struct {
	// valid values are: `NONE` or `COUNT`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
