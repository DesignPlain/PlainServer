package types

type Alb_ListenerRuleActionForward struct {
	// The target group stickiness for the rule.
	Stickiness Alb_ListenerRuleActionForwardStickiness `json:"stickiness,omitempty" yaml:"stickiness,omitempty"`

	// One or more target groups block.
	TargetGroups []Alb_ListenerRuleActionForwardTargetGroup `json:"targetGroups,omitempty" yaml:"targetGroups,omitempty"`
}
