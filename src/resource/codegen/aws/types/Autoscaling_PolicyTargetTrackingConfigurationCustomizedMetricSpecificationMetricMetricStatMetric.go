package types

type Autoscaling_PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetric struct {
	// Dimensions of the metric.
	Dimensions []Autoscaling_PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStatMetricDimension `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`

	// Name of the metric.
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`

	// Namespace of the metric.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}
