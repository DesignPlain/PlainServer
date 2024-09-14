package types

type Cloudwatch_LogMetricFilterMetricTransformation struct {
	// The value to emit when a filter pattern does not match a log event. Conflicts with `dimensions`.
	DefaultValue string `json:"defaultValue,omitempty" yaml:"defaultValue,omitempty"`

	// Map of fields to use as dimensions for the metric. Up to 3 dimensions are allowed. Conflicts with `default_value`.
	Dimensions map[string]string `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`

	// The name of the CloudWatch metric to which the monitored log information should be published (e.g., `ErrorCount`)
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The destination namespace of the CloudWatch metric.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	// The unit to assign to the metric. If you omit this, the unit is set as `None`.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	// What to publish to the metric. For example, if you're counting the occurrences of a particular term like "Error", the value will be "1" for each occurrence. If you're counting the bytes transferred the published value will be the value in the log event.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
