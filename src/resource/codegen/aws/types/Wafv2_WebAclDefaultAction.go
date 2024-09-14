package types

type Wafv2_WebAclDefaultAction struct {
	// Specifies that AWS WAF should allow requests by default. See `allow` below for details.
	Allow Wafv2_WebAclDefaultActionAllow `json:"allow,omitempty" yaml:"allow,omitempty"`

	// Specifies that AWS WAF should block requests by default. See `block` below for details.
	Block Wafv2_WebAclDefaultActionBlock `json:"block,omitempty" yaml:"block,omitempty"`
}
