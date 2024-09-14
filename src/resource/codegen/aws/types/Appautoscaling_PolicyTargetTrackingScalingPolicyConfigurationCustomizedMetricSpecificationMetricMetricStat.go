package types

type Appautoscaling_PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStat struct {
	// Structure that defines the CloudWatch metric to return, including the metric name, namespace, and dimensions.
	Metric Appautoscaling_PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricMetricStatMetric `json:"metric,omitempty" yaml:"metric,omitempty"`

	// Statistic of the metrics to return.
	Stat string `json:"stat,omitempty" yaml:"stat,omitempty"`

	// Unit of the metrics to return.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`
}
