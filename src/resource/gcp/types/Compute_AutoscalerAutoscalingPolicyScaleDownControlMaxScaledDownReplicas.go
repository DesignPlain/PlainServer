package types

type Compute_AutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicas struct {
	/*
	   Specifies a fixed number of VM instances. This must be a positive
	   integer.
	*/
	Fixed int `json:"fixed,omitempty" yaml:"fixed,omitempty"`

	/*
	   Specifies a percentage of instances between 0 to 100%!!(MISSING),(MISSING) inclusive.
	   For example, specify 80 for 80%!!(MISSING)
	   (MISSING)
	*/
	Percent int `json:"percent,omitempty" yaml:"percent,omitempty"`
}
