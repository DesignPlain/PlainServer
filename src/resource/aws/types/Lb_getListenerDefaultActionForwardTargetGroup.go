package types

type Lb_getListenerDefaultActionForwardTargetGroup struct {
	// ARN of the listener. Required if `load_balancer_arn` and `port` is not set.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	//
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}
