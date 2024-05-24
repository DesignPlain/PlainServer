package types

type Vpclattice_ListenerRuleAction struct {
	// Describes the rule action that returns a custom HTTP response.
	FixedResponse Vpclattice_ListenerRuleActionFixedResponse `json:"fixedResponse,omitempty" yaml:"fixedResponse,omitempty"`

	// The forward action. Traffic that matches the rule is forwarded to the specified target groups.
	Forward Vpclattice_ListenerRuleActionForward `json:"forward,omitempty" yaml:"forward,omitempty"`
}
