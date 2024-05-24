package types

type Waf_WebAclRuleOverrideAction struct {
	/*
	   Specifies how you want AWS WAF to respond to requests that don't match the criteria in any of the `rules`.
	   e.g., `ALLOW` or `BLOCK`
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
