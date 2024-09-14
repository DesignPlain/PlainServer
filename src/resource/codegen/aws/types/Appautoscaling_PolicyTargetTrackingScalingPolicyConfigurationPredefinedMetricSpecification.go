package types

type Appautoscaling_PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification struct {
	// Metric type.
	PredefinedMetricType string `json:"predefinedMetricType,omitempty" yaml:"predefinedMetricType,omitempty"`

	// Reserved for future use if the `predefined_metric_type` is not `ALBRequestCountPerTarget`. If the `predefined_metric_type` is `ALBRequestCountPerTarget`, you must specify this argument. Documentation can be found at: [AWS Predefined Scaling Metric Specification](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_PredefinedScalingMetricSpecification.html). Must be less than or equal to 1023 characters in length.
	ResourceLabel string `json:"resourceLabel,omitempty" yaml:"resourceLabel,omitempty"`
}
