package types

type Autoscaling_PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecification struct {
	// List of up to 10 structures that defines custom load metric in predictive scaling policy
	MetricDataQueries []Autoscaling_PolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecificationMetricDataQuery `json:"metricDataQueries,omitempty" yaml:"metricDataQueries,omitempty"`
}
