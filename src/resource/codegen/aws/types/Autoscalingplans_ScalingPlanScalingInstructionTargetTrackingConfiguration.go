package types

type Autoscalingplans_ScalingPlanScalingInstructionTargetTrackingConfiguration struct {
	/*
	   Customized metric. You can specify either `customized_scaling_metric_specification` or `predefined_scaling_metric_specification`.
	   More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_CustomizedScalingMetricSpecification.html).
	*/
	CustomizedScalingMetricSpecification Autoscalingplans_ScalingPlanScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecification `json:"customizedScalingMetricSpecification,omitempty" yaml:"customizedScalingMetricSpecification,omitempty"`

	// Boolean indicating whether scale in by the target tracking scaling policy is disabled. Defaults to `false`.
	DisableScaleIn bool `json:"disableScaleIn,omitempty" yaml:"disableScaleIn,omitempty"`

	/*
	   Estimated time, in seconds, until a newly launched instance can contribute to the CloudWatch metrics.
	   This value is used only if the resource is an Auto Scaling group.
	*/
	EstimatedInstanceWarmup int `json:"estimatedInstanceWarmup,omitempty" yaml:"estimatedInstanceWarmup,omitempty"`

	/*
	   Predefined metric. You can specify either `predefined_scaling_metric_specification` or `customized_scaling_metric_specification`.
	   More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_PredefinedScalingMetricSpecification.html).
	*/
	PredefinedScalingMetricSpecification Autoscalingplans_ScalingPlanScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification `json:"predefinedScalingMetricSpecification,omitempty" yaml:"predefinedScalingMetricSpecification,omitempty"`

	/*
	   Amount of time, in seconds, after a scale in activity completes before another scale in activity can start.
	   This value is not used if the scalable resource is an Auto Scaling group.
	*/
	ScaleInCooldown int `json:"scaleInCooldown,omitempty" yaml:"scaleInCooldown,omitempty"`

	/*
	   Amount of time, in seconds, after a scale-out activity completes before another scale-out activity can start.
	   This value is not used if the scalable resource is an Auto Scaling group.
	*/
	ScaleOutCooldown int `json:"scaleOutCooldown,omitempty" yaml:"scaleOutCooldown,omitempty"`

	// Target value for the metric.
	TargetValue float64 `json:"targetValue,omitempty" yaml:"targetValue,omitempty"`
}
