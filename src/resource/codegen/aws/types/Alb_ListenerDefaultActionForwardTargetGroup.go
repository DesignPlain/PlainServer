package types

type Alb_ListenerDefaultActionForwardTargetGroup struct {
	// Weight. The range is 0 to 999.
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`

	/*
	   ARN of the target group.

	   The following arguments are optional:
	*/
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`
}
