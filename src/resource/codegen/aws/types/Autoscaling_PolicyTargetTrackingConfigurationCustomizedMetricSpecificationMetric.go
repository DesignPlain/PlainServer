package types

type Autoscaling_PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetric struct {
	// Structure that defines CloudWatch metric to be used in target tracking scaling policy. You must specify either `expression` or `metric_stat`, but not both.
	MetricStat Autoscaling_PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricMetricStat `json:"metricStat,omitempty" yaml:"metricStat,omitempty"`

	// Boolean that indicates whether to return the timestamps and raw data values of this metric, the default is true
	ReturnData bool `json:"returnData,omitempty" yaml:"returnData,omitempty"`

	// Math expression used on the returned metric. You must specify either `expression` or `metric_stat`, but not both.
	Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`

	// Short name for the metric used in target tracking scaling policy.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Human-readable label for this metric or expression.
	Label string `json:"label,omitempty" yaml:"label,omitempty"`
}
