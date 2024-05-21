package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type FirewallPolicyRule struct {
	// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
	TargetResources []string `json:"targetResources,omitempty" yaml:"targetResources,omitempty"`

	// A list of service accounts indicating the sets of instances that are applied with this rule.
	TargetServiceAccounts []string `json:"targetServiceAccounts,omitempty" yaml:"targetServiceAccounts,omitempty"`

	// The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "goto_next".
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// An optional description for this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	Match types.Compute_FirewallPolicyRuleMatch `json:"match,omitempty" yaml:"match,omitempty"`

	// The direction in which this rule applies. Possible values: INGRESS, EGRESS
	Direction string `json:"direction,omitempty" yaml:"direction,omitempty"`

	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.
	EnableLogging bool `json:"enableLogging,omitempty" yaml:"enableLogging,omitempty"`

	// The firewall policy of the resource.
	FirewallPolicy string `json:"firewallPolicy,omitempty" yaml:"firewallPolicy,omitempty"`
}
