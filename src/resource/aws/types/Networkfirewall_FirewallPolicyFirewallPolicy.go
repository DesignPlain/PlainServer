package types

type Networkfirewall_FirewallPolicyFirewallPolicy struct {
	// Set of configuration blocks containing references to the stateful rule groups that are used in the policy. See Stateful Rule Group Reference below for details.
	StatefulRuleGroupReferences []Networkfirewall_FirewallPolicyFirewallPolicyStatefulRuleGroupReference `json:"statefulRuleGroupReferences,omitempty" yaml:"statefulRuleGroupReferences,omitempty"`

	// Set of configuration blocks containing references to the stateless rule groups that are used in the policy. See Stateless Rule Group Reference below for details.
	StatelessRuleGroupReferences []Networkfirewall_FirewallPolicyFirewallPolicyStatelessRuleGroupReference `json:"statelessRuleGroupReferences,omitempty" yaml:"statelessRuleGroupReferences,omitempty"`

	// The (ARN) of the TLS Inspection policy to attach to the FW Policy.  This must be added at creation of the resource per AWS documentation. "You can only add a TLS inspection configuration to a new policy, not to an existing policy."  This cannot be removed from a FW Policy.
	TlsInspectionConfigurationArn string `json:"tlsInspectionConfigurationArn,omitempty" yaml:"tlsInspectionConfigurationArn,omitempty"`

	// . Contains variables that you can use to override default Suricata settings in your firewall policy. See Rule Variables for details.
	PolicyVariables Networkfirewall_FirewallPolicyFirewallPolicyPolicyVariables `json:"policyVariables,omitempty" yaml:"policyVariables,omitempty"`

	// Set of actions to take on a packet if it does not match any stateful rules in the policy. This can only be specified if the policy has a `stateful_engine_options` block with a `rule_order` value of `STRICT_ORDER`. You can specify one of either or neither values of `aws:drop_strict` or `aws:drop_established`, as well as any combination of `aws:alert_strict` and `aws:alert_established`.
	StatefulDefaultActions []string `json:"statefulDefaultActions,omitempty" yaml:"statefulDefaultActions,omitempty"`

	// A configuration block that defines options on how the policy handles stateful rules. See Stateful Engine Options below for details.
	StatefulEngineOptions Networkfirewall_FirewallPolicyFirewallPolicyStatefulEngineOptions `json:"statefulEngineOptions,omitempty" yaml:"statefulEngineOptions,omitempty"`

	// Set of configuration blocks describing the custom action definitions that are available for use in the firewall policy's `stateless_default_actions`. See Stateless Custom Action below for details.
	StatelessCustomActions []Networkfirewall_FirewallPolicyFirewallPolicyStatelessCustomAction `json:"statelessCustomActions,omitempty" yaml:"statelessCustomActions,omitempty"`

	/*
	   Set of actions to take on a packet if it does not match any of the stateless rules in the policy. You must specify one of the standard actions including: `aws:drop`, `aws:pass`, or `aws:forward_to_sfe`.
	   In addition, you can specify custom actions that are compatible with your standard action choice. If you want non-matching packets to be forwarded for stateful inspection, specify `aws:forward_to_sfe`.
	*/
	StatelessDefaultActions []string `json:"statelessDefaultActions,omitempty" yaml:"statelessDefaultActions,omitempty"`

	/*
	   Set of actions to take on a fragmented packet if it does not match any of the stateless rules in the policy. You must specify one of the standard actions including: `aws:drop`, `aws:pass`, or `aws:forward_to_sfe`.
	   In addition, you can specify custom actions that are compatible with your standard action choice. If you want non-matching packets to be forwarded for stateful inspection, specify `aws:forward_to_sfe`.
	*/
	StatelessFragmentDefaultActions []string `json:"statelessFragmentDefaultActions,omitempty" yaml:"statelessFragmentDefaultActions,omitempty"`
}
