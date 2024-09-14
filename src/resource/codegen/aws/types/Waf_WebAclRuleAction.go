package types

type Waf_WebAclRuleAction struct {
	// valid values are: `BLOCK`, `ALLOW`, or `COUNT`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
