package types

type Wafv2_RuleGroupRuleStatementIpSetReferenceStatement struct {
	// The Amazon Resource Name (ARN) of the IP Set that this statement references.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin. See IPSet Forwarded IP Config below for more details.
	IpSetForwardedIpConfig Wafv2_RuleGroupRuleStatementIpSetReferenceStatementIpSetForwardedIpConfig `json:"ipSetForwardedIpConfig,omitempty" yaml:"ipSetForwardedIpConfig,omitempty"`
}
