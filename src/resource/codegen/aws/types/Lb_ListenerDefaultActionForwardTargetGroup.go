package types

type Lb_ListenerDefaultActionForwardTargetGroup struct {
	/*
	   ARN of the target group.

	   The following arguments are optional:
	*/
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// Weight. The range is 0 to 999.
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}
