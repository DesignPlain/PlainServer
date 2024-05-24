package types

type Autoscaling_PolicyTargetTrackingConfigurationCustomizedMetricSpecification struct {
	// Dimensions of the metric.
	MetricDimensions []Autoscaling_PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension `json:"metricDimensions,omitempty" yaml:"metricDimensions,omitempty"`

	// Name of the metric.
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`

	// Metrics to include, as a metric data query.
	Metrics []Autoscaling_PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetric `json:"metrics,omitempty" yaml:"metrics,omitempty"`

	// Namespace of the metric.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	// Statistic of the metric.
	Statistic string `json:"statistic,omitempty" yaml:"statistic,omitempty"`

	// Unit of the metric.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`
}
