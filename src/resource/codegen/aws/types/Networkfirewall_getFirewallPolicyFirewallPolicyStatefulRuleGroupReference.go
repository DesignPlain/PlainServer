package types

type Networkfirewall_getFirewallPolicyFirewallPolicyStatefulRuleGroupReference struct {
	//
	Overrides []Networkfirewall_getFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceOverride `json:"overrides,omitempty" yaml:"overrides,omitempty"`

	//
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	//
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`
}
