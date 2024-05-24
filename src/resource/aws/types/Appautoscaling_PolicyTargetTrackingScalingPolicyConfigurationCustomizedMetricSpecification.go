package types

type Appautoscaling_PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification struct {
	// Statistic of the metric. Valid values: `Average`, `Minimum`, `Maximum`, `SampleCount`, and `Sum`.
	Statistic string `json:"statistic,omitempty" yaml:"statistic,omitempty"`

	// Unit of the metrics to return.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	// Dimensions of the metric.
	Dimensions []Appautoscaling_PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`

	// Name of the metric.
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`

	// Metrics to include, as a metric data query.
	Metrics []Appautoscaling_PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetric `json:"metrics,omitempty" yaml:"metrics,omitempty"`

	// Namespace of the metric.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}
