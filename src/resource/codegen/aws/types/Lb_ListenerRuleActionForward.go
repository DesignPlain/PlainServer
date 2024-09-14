package types

type Lb_ListenerRuleActionForward struct {
	// The target group stickiness for the rule.
	Stickiness Lb_ListenerRuleActionForwardStickiness `json:"stickiness,omitempty" yaml:"stickiness,omitempty"`

	// One or more target groups block.
	TargetGroups []Lb_ListenerRuleActionForwardTargetGroup `json:"targetGroups,omitempty" yaml:"targetGroups,omitempty"`
}
