package types

type Networkfirewall_FirewallPolicyFirewallPolicyStatefulRuleGroupReference struct {
	// Configuration block for override values
	Override Networkfirewall_FirewallPolicyFirewallPolicyStatefulRuleGroupReferenceOverride `json:"override,omitempty" yaml:"override,omitempty"`

	// An integer setting that indicates the order in which to apply the stateful rule groups in a single policy. This argument must be specified if the policy has a `stateful_engine_options` block with a `rule_order` value of `STRICT_ORDER`. AWS Network Firewall applies each stateful rule group to a packet starting with the group that has the lowest priority setting.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// The Amazon Resource Name (ARN) of the stateful rule group.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`
}
