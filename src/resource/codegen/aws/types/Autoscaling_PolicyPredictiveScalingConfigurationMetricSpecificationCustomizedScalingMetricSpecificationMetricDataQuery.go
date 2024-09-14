package types

type Autoscaling_PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQuery struct {
	// Math expression used on the returned metric. You must specify either `expression` or `metric_stat`, but not both.
	Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`

	// Short name for the metric used in predictive scaling policy.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Human-readable label for this metric or expression.
	Label string `json:"label,omitempty" yaml:"label,omitempty"`

	// Structure that defines CloudWatch metric to be used in predictive scaling policy. You must specify either `expression` or `metric_stat`, but not both.
	MetricStat Autoscaling_PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQueryMetricStat `json:"metricStat,omitempty" yaml:"metricStat,omitempty"`

	// Boolean that indicates whether to return the timestamps and raw data values of this metric, the default is true
	ReturnData bool `json:"returnData,omitempty" yaml:"returnData,omitempty"`
}
