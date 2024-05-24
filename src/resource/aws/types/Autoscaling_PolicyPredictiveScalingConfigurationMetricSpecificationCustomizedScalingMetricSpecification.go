package types

type Autoscaling_PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecification struct {
	// List of up to 10 structures that defines custom scaling metric in predictive scaling policy
	MetricDataQueries []Autoscaling_PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedScalingMetricSpecificationMetricDataQuery `json:"metricDataQueries,omitempty" yaml:"metricDataQueries,omitempty"`
}
