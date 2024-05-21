package types

type Compute_RegionAutoscalerAutoscalingPolicy struct {
	/*
	   The maximum number of instances that the autoscaler can scale up
	   to. This is required when creating or updating an autoscaler. The
	   maximum number of replicas should not be lower than minimal number
	   of replicas.
	*/
	MaxReplicas int `json:"maxReplicas,omitempty" yaml:"maxReplicas,omitempty"`

	/*
	   The minimum number of replicas that the autoscaler can scale down
	   to. This cannot be less than 0. If not provided, autoscaler will
	   choose a default value depending on maximum number of instances
	   allowed.
	*/
	MinReplicas int `json:"minReplicas,omitempty" yaml:"minReplicas,omitempty"`

	// Defines operating mode for this policy.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	/*
	   Defines scale down controls to reduce the risk of response latency
	   and outages due to abrupt scale-in events
	   Structure is documented below.
	*/
	ScaleDownControl Compute_RegionAutoscalerAutoscalingPolicyScaleDownControl `json:"scaleDownControl,omitempty" yaml:"scaleDownControl,omitempty"`

	/*
	   Defines scale in controls to reduce the risk of response latency
	   and outages due to abrupt scale-in events
	   Structure is documented below.
	*/
	ScaleInControl Compute_RegionAutoscalerAutoscalingPolicyScaleInControl `json:"scaleInControl,omitempty" yaml:"scaleInControl,omitempty"`

	/*
	   Configuration parameters of autoscaling based on a load balancer.
	   Structure is documented below.
	*/
	LoadBalancingUtilization Compute_RegionAutoscalerAutoscalingPolicyLoadBalancingUtilization `json:"loadBalancingUtilization,omitempty" yaml:"loadBalancingUtilization,omitempty"`

	/*
	   Defines the CPU utilization policy that allows the autoscaler to
	   scale based on the average CPU utilization of a managed instance
	   group.
	   Structure is documented below.
	*/
	CpuUtilization Compute_RegionAutoscalerAutoscalingPolicyCpuUtilization `json:"cpuUtilization,omitempty" yaml:"cpuUtilization,omitempty"`

	/*
	   Configuration parameters of autoscaling based on a custom metric.
	   Structure is documented below.
	*/
	Metrics []Compute_RegionAutoscalerAutoscalingPolicyMetric `json:"metrics,omitempty" yaml:"metrics,omitempty"`

	/*
	   Scaling schedules defined for an autoscaler. Multiple schedules can be set on an autoscaler and they can overlap.
	   Structure is documented below.
	*/
	ScalingSchedules []Compute_RegionAutoscalerAutoscalingPolicyScalingSchedule `json:"scalingSchedules,omitempty" yaml:"scalingSchedules,omitempty"`

	/*
	   The number of seconds that the autoscaler should wait before it
	   starts collecting information from a new instance. This prevents
	   the autoscaler from collecting information when the instance is
	   initializing, during which the collected usage would not be
	   reliable. The default time autoscaler waits is 60 seconds.
	   Virtual machine initialization times might vary because of
	   numerous factors. We recommend that you test how long an
	   instance may take to initialize. To do this, create an instance
	   and time the startup process.
	*/
	CooldownPeriod int `json:"cooldownPeriod,omitempty" yaml:"cooldownPeriod,omitempty"`
}
