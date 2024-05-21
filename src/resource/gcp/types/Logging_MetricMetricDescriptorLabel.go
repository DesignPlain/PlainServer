package types

type Logging_MetricMetricDescriptorLabel struct {
	// A human-readable description for the label.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The label key.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	/*
	   The type of data that can be assigned to the label.
	   Default value is `STRING`.
	   Possible values are: `BOOL`, `INT64`, `STRING`.
	*/
	ValueType string `json:"valueType,omitempty" yaml:"valueType,omitempty"`
}
