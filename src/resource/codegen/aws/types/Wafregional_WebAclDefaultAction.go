package types

type Wafregional_WebAclDefaultAction struct {
	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a ruleE.g., `ALLOW`, `BLOCK` or `COUNT`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
