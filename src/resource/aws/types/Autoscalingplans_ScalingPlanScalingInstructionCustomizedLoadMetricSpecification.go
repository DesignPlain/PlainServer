package types

type Autoscalingplans_ScalingPlanScalingInstructionCustomizedLoadMetricSpecification struct {
	// Dimensions of the metric.
	Dimensions map[string]string `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`

	// Name of the metric.
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`

	// Namespace of the metric.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	// Statistic of the metric. Currently, the value must always be `Sum`.
	Statistic string `json:"statistic,omitempty" yaml:"statistic,omitempty"`

	// Unit of the metric.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`
}
