package types

type Compute_AutoscalerAutoscalingPolicyLoadBalancingUtilization struct {
	/*
	   Fraction of backend capacity utilization (set in HTTP(s) load
	   balancing configuration) that autoscaler should maintain. Must
	   be a positive float value. If not defined, the default is 0.8.
	*/
	Target float64 `json:"target,omitempty" yaml:"target,omitempty"`
}
