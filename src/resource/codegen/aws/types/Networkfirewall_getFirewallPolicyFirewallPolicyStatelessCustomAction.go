package types

type Networkfirewall_getFirewallPolicyFirewallPolicyStatelessCustomAction struct {
	//
	ActionName string `json:"actionName,omitempty" yaml:"actionName,omitempty"`

	//
	ActionDefinitions []Networkfirewall_getFirewallPolicyFirewallPolicyStatelessCustomActionActionDefinition `json:"actionDefinitions,omitempty" yaml:"actionDefinitions,omitempty"`
}
