package compute

import types "libds/gcp/types"

type OrganizationSecurityPolicyRule struct {
	// If set to true, the specified action is not enforced.
	Preview bool `json:"preview,omitempty" yaml:"preview,omitempty"`

	/*
	   An integer indicating the priority of a rule in the list. The priority must be a value
	   between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
	   highest priority and 2147483647 is the lowest prority.
	*/
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	/*
	   A list of network resource URLs to which this rule applies.
	   This field allows you to control which network's VMs get
	   this rule. If this field is left blank, all VMs
	   within the organization will receive the rule.
	*/
	TargetResources []string `json:"targetResources,omitempty" yaml:"targetResources,omitempty"`

	/*
	   The Action to perform when the client connection triggers the rule. Can currently be either
	   "allow", "deny" or "goto_next".
	*/
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// A description of the rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Denotes whether to enable logging for a particular rule.
	   If logging is enabled, logs will be exported to the
	   configured export destination in Stackdriver.
	*/
	EnableLogging bool `json:"enableLogging,omitempty" yaml:"enableLogging,omitempty"`

	// The ID of the OrganizationSecurityPolicy this rule applies to.
	PolicyId string `json:"policyId,omitempty" yaml:"policyId,omitempty"`

	/*
	   The direction in which this rule applies. If unspecified an INGRESS rule is created.
	   Possible values are: `INGRESS`, `EGRESS`.
	*/
	Direction string `json:"direction,omitempty" yaml:"direction,omitempty"`

	/*
	   A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	   Structure is documented below.
	*/
	Match types.Compute_OrganizationSecurityPolicyRuleMatch `json:"match,omitempty" yaml:"match,omitempty"`

	/*
	   A list of service accounts indicating the sets of
	   instances that are applied with this rule.
	*/
	TargetServiceAccounts []string `json:"targetServiceAccounts,omitempty" yaml:"targetServiceAccounts,omitempty"`
}
