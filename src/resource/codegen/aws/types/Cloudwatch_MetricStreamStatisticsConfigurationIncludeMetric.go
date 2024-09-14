package types

type Cloudwatch_MetricStreamStatisticsConfigurationIncludeMetric struct {
	// The name of the metric.
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`

	//
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}
