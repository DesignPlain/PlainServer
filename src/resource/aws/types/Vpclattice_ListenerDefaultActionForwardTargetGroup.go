package types

type Vpclattice_ListenerDefaultActionForwardTargetGroup struct {
	// ID or Amazon Resource Name (ARN) of the target group.
	TargetGroupIdentifier string `json:"targetGroupIdentifier,omitempty" yaml:"targetGroupIdentifier,omitempty"`

	/*
	   Determines how requests are distributed to the target group. Only required if you specify multiple target groups for a forward action. For example, if you specify two target groups, one with a
	   weight of 10 and the other with a weight of 20, the target group with a weight of 20 receives twice as many requests as the other target group. See [Listener rules](https://docs.aws.amazon.com/vpc-lattice/latest/ug/listeners.html#listener-rules) in the AWS documentation for additional examples. Default: `100`.
	*/
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}
