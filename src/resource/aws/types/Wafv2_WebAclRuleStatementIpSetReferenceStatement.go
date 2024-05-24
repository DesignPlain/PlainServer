package types

type Wafv2_WebAclRuleStatementIpSetReferenceStatement struct {
	// The Amazon Resource Name (ARN) of the IP Set that this statement references.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// Configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin. See `ip_set_forwarded_ip_config` below for more details.
	IpSetForwardedIpConfig Wafv2_WebAclRuleStatementIpSetReferenceStatementIpSetForwardedIpConfig `json:"ipSetForwardedIpConfig,omitempty" yaml:"ipSetForwardedIpConfig,omitempty"`
}
