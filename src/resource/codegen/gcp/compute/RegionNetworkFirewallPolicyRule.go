package compute

import types "libds/gcp/types"

type RegionNetworkFirewallPolicyRule struct {
	// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// An optional name for the rule. This field is not a unique identifier and can be updated.
	RuleName string `json:"ruleName,omitempty" yaml:"ruleName,omitempty"`

	// A list of secure tags that controls which instances the firewall rule applies to. If <code>targetSecureTag</code> are specified, then the firewall rule applies only to instances in the VPC network that have one of those EFFECTIVE secure tags, if all the target_secure_tag are in INEFFECTIVE state, then this rule will be ignored. <code>targetSecureTag</code> may not be set at the same time as <code>targetServiceAccounts</code>. If neither <code>targetServiceAccounts</code> nor <code>targetSecureTag</code> are specified, the firewall rule applies to all instances on the specified network. Maximum number of target label tags allowed is 256.
	TargetSecureTags []types.Compute_RegionNetworkFirewallPolicyRuleTargetSecureTag `json:"targetSecureTags,omitempty" yaml:"targetSecureTags,omitempty"`

	// A list of service accounts indicating the sets of instances that are applied with this rule.
	TargetServiceAccounts []string `json:"targetServiceAccounts,omitempty" yaml:"targetServiceAccounts,omitempty"`

	// The direction in which this rule applies. Possible values: INGRESS, EGRESS
	Direction string `json:"direction,omitempty" yaml:"direction,omitempty"`

	// The firewall policy of the resource.
	FirewallPolicy string `json:"firewallPolicy,omitempty" yaml:"firewallPolicy,omitempty"`

	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.
	EnableLogging bool `json:"enableLogging,omitempty" yaml:"enableLogging,omitempty"`

	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	Match types.Compute_RegionNetworkFirewallPolicyRuleMatch `json:"match,omitempty" yaml:"match,omitempty"`

	// The location of this resource.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "goto_next".
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// An optional description for this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
