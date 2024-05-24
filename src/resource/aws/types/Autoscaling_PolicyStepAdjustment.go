package types

type Autoscaling_PolicyStepAdjustment struct {
	/*
	   Upper bound for the
	   difference between the alarm threshold and the CloudWatch metric.
	   Without a value, AWS will treat this bound as positive infinity. The upper bound
	   must be greater than the lower bound.

	   Notice the bounds are --relative-- to the alarm threshold, meaning that the starting point is not 0%!,(MISSING) but the alarm threshold. Check the official [docs](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html#as-scaling-steps) for a detailed example.

	   The following arguments are only available to "TargetTrackingScaling" type policies:
	*/
	MetricIntervalUpperBound string `json:"metricIntervalUpperBound,omitempty" yaml:"metricIntervalUpperBound,omitempty"`

	/*
	   Number of members by which to
	   scale, when the adjustment bounds are breached. A positive value scales
	   up. A negative value scales down.
	*/
	ScalingAdjustment int `json:"scalingAdjustment,omitempty" yaml:"scalingAdjustment,omitempty"`

	/*
	   Lower bound for the
	   difference between the alarm threshold and the CloudWatch metric.
	   Without a value, AWS will treat this bound as negative infinity.
	*/
	MetricIntervalLowerBound string `json:"metricIntervalLowerBound,omitempty" yaml:"metricIntervalLowerBound,omitempty"`
}
