package types

type Networkfirewall_FirewallPolicyFirewallPolicyStatefulEngineOptions struct {
	// Indicates how to manage the order of stateful rule evaluation for the policy. Default value: `DEFAULT_ACTION_ORDER`. Valid values: `DEFAULT_ACTION_ORDER`, `STRICT_ORDER`.
	RuleOrder string `json:"ruleOrder,omitempty" yaml:"ruleOrder,omitempty"`

	// Describes how to treat traffic which has broken midstream. Default value: `DROP`. Valid values: `DROP`, `CONTINUE`, `REJECT`.
	StreamExceptionPolicy string `json:"streamExceptionPolicy,omitempty" yaml:"streamExceptionPolicy,omitempty"`
}
