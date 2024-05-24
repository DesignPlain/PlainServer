package types

type Networkfirewall_FirewallPolicyFirewallPolicyStatelessRuleGroupReference struct {
	// An integer setting that indicates the order in which to run the stateless rule groups in a single policy. AWS Network Firewall applies each stateless rule group to a packet starting with the group that has the lowest priority setting.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// The Amazon Resource Name (ARN) of the stateless rule group.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`
}
