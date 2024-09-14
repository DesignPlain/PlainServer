package types

type Appautoscaling_PolicyTargetTrackingScalingPolicyConfiguration struct {
	// Custom CloudWatch metric. Documentation can be found  at: [AWS Customized Metric Specification](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CustomizedMetricSpecification.html). See supported fields below.
	CustomizedMetricSpecification Appautoscaling_PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification `json:"customizedMetricSpecification,omitempty" yaml:"customizedMetricSpecification,omitempty"`

	// Whether scale in by the target tracking policy is disabled. If the value is true, scale in is disabled and the target tracking policy won't remove capacity from the scalable resource. Otherwise, scale in is enabled and the target tracking policy can remove capacity from the scalable resource. The default value is `false`.
	DisableScaleIn bool `json:"disableScaleIn,omitempty" yaml:"disableScaleIn,omitempty"`

	// Predefined metric. See supported fields below.
	PredefinedMetricSpecification Appautoscaling_PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification `json:"predefinedMetricSpecification,omitempty" yaml:"predefinedMetricSpecification,omitempty"`

	// Amount of time, in seconds, after a scale in activity completes before another scale in activity can start.
	ScaleInCooldown int `json:"scaleInCooldown,omitempty" yaml:"scaleInCooldown,omitempty"`

	// Amount of time, in seconds, after a scale out activity completes before another scale out activity can start.
	ScaleOutCooldown int `json:"scaleOutCooldown,omitempty" yaml:"scaleOutCooldown,omitempty"`

	// Target value for the metric.
	TargetValue float64 `json:"targetValue,omitempty" yaml:"targetValue,omitempty"`
}
