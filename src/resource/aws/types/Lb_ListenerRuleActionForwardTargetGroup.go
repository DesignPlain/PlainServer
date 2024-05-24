package types

type Lb_ListenerRuleActionForwardTargetGroup struct {
	// The Amazon Resource Name (ARN) of the target group.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// The weight. The range is 0 to 999.
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}
