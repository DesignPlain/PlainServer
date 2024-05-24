package types

type Autoscaling_PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecification struct {
	// List of up to 10 structures that defines custom capacity metric in predictive scaling policy
	MetricDataQueries []Autoscaling_PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQuery `json:"metricDataQueries,omitempty" yaml:"metricDataQueries,omitempty"`
}
