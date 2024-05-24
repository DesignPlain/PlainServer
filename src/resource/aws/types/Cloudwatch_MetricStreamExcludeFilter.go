package types

type Cloudwatch_MetricStreamExcludeFilter struct {
	// An array that defines the metrics you want to exclude for this metric namespace
	MetricNames []string `json:"metricNames,omitempty" yaml:"metricNames,omitempty"`

	// Name of the metric namespace in the filter.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}
