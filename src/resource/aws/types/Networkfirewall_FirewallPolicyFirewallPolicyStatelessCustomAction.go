package types

type Networkfirewall_FirewallPolicyFirewallPolicyStatelessCustomAction struct {
	// A configuration block describing the custom action associated with the `action_name`. See Action Definition below for details.
	ActionDefinition Networkfirewall_FirewallPolicyFirewallPolicyStatelessCustomActionActionDefinition `json:"actionDefinition,omitempty" yaml:"actionDefinition,omitempty"`

	// A friendly name of the custom action.
	ActionName string `json:"actionName,omitempty" yaml:"actionName,omitempty"`
}
