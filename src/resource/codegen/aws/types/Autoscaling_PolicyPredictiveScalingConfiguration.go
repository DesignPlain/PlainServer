package types

type Autoscaling_PolicyPredictiveScalingConfiguration struct {
	// Defines the behavior that should be applied if the forecast capacity approaches or exceeds the maximum capacity of the Auto Scaling group. Valid values are `HonorMaxCapacity` or `IncreaseMaxCapacity`. Default is `HonorMaxCapacity`.
	MaxCapacityBreachBehavior string `json:"maxCapacityBreachBehavior,omitempty" yaml:"maxCapacityBreachBehavior,omitempty"`

	// Size of the capacity buffer to use when the forecast capacity is close to or exceeds the maximum capacity. Valid range is `0` to `100`. If set to `0`, Amazon EC2 Auto Scaling may scale capacity higher than the maximum capacity to equal but not exceed forecast capacity.
	MaxCapacityBuffer string `json:"maxCapacityBuffer,omitempty" yaml:"maxCapacityBuffer,omitempty"`

	// This structure includes the metrics and target utilization to use for predictive scaling.
	MetricSpecification Autoscaling_PolicyPredictiveScalingConfigurationMetricSpecification `json:"metricSpecification,omitempty" yaml:"metricSpecification,omitempty"`

	// Predictive scaling mode. Valid values are `ForecastAndScale` and `ForecastOnly`. Default is `ForecastOnly`.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	// Amount of time, in seconds, by which the instance launch time can be advanced. Minimum is `0`.
	SchedulingBufferTime string `json:"schedulingBufferTime,omitempty" yaml:"schedulingBufferTime,omitempty"`
}
