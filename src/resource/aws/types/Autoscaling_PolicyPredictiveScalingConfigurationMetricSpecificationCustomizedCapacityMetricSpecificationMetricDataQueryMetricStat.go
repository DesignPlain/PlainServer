package types

type Autoscaling_PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStat struct {
	// Unit of the metrics to return.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	// Structure that defines the CloudWatch metric to return, including the metric name, namespace, and dimensions.
	Metric Autoscaling_PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatMetric `json:"metric,omitempty" yaml:"metric,omitempty"`

	// Statistic of the metrics to return.
	Stat string `json:"stat,omitempty" yaml:"stat,omitempty"`
}
