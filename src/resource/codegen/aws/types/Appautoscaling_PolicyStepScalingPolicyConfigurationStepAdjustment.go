package types

type Appautoscaling_PolicyStepScalingPolicyConfigurationStepAdjustment struct {
	// Lower bound for the difference between the alarm threshold and the CloudWatch metric. Without a value, AWS will treat this bound as negative infinity.
	MetricIntervalLowerBound string `json:"metricIntervalLowerBound,omitempty" yaml:"metricIntervalLowerBound,omitempty"`

	// Upper bound for the difference between the alarm threshold and the CloudWatch metric. Without a value, AWS will treat this bound as infinity. The upper bound must be greater than the lower bound.
	MetricIntervalUpperBound string `json:"metricIntervalUpperBound,omitempty" yaml:"metricIntervalUpperBound,omitempty"`

	// Number of members by which to scale, when the adjustment bounds are breached. A positive value scales up. A negative value scales down.
	ScalingAdjustment int `json:"scalingAdjustment,omitempty" yaml:"scalingAdjustment,omitempty"`
}
