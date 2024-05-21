package types

type Spanner_InstanceAutoscalingConfig struct {
	/*
	   Defines scale in controls to reduce the risk of response latency
	   and outages due to abrupt scale-in events. Users can define the minimum and
	   maximum compute capacity allocated to the instance, and the autoscaler will
	   only scale within that range. Users can either use nodes or processing
	   units to specify the limits, but should use the same unit to set both the
	   min_limit and max_limit.
	   Structure is documented below.
	*/
	AutoscalingLimits Spanner_InstanceAutoscalingConfigAutoscalingLimits `json:"autoscalingLimits,omitempty" yaml:"autoscalingLimits,omitempty"`

	/*
	   Defines scale in controls to reduce the risk of response latency
	   and outages due to abrupt scale-in events
	   Structure is documented below.
	*/
	AutoscalingTargets Spanner_InstanceAutoscalingConfigAutoscalingTargets `json:"autoscalingTargets,omitempty" yaml:"autoscalingTargets,omitempty"`
}
