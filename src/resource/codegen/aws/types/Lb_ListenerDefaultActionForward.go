package types

type Lb_ListenerDefaultActionForward struct {
	// Configuration block for target group stickiness for the rule. Detailed below.
	Stickiness Lb_ListenerDefaultActionForwardStickiness `json:"stickiness,omitempty" yaml:"stickiness,omitempty"`

	/*
	   Set of 1-5 target group blocks. Detailed below.

	   The following arguments are optional:
	*/
	TargetGroups []Lb_ListenerDefaultActionForwardTargetGroup `json:"targetGroups,omitempty" yaml:"targetGroups,omitempty"`
}
