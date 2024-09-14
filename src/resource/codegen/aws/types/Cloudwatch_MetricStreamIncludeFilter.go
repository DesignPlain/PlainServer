package types

type Cloudwatch_MetricStreamIncludeFilter struct {
	// An array that defines the metrics you want to include for this metric namespace
	MetricNames []string `json:"metricNames,omitempty" yaml:"metricNames,omitempty"`

	// Name of the metric namespace in the filter.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}
