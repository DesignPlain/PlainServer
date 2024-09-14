package types

type Autoscaling_PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetric struct {
	// Namespace of the metric.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	// Dimensions of the metric.
	Dimensions []Autoscaling_PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStatMetricDimension `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`

	// Name of the metric.
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`
}
