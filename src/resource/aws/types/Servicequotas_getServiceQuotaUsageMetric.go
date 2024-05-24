package types

type Servicequotas_getServiceQuotaUsageMetric struct {
	// The metric dimensions.
	MetricDimensions []Servicequotas_getServiceQuotaUsageMetricMetricDimension `json:"metricDimensions,omitempty" yaml:"metricDimensions,omitempty"`

	// The name of the metric.
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`

	// The namespace of the metric.
	MetricNamespace string `json:"metricNamespace,omitempty" yaml:"metricNamespace,omitempty"`

	// The metric statistic that AWS recommend you use when determining quota usage.
	MetricStatisticRecommendation string `json:"metricStatisticRecommendation,omitempty" yaml:"metricStatisticRecommendation,omitempty"`
}
