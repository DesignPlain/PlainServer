package types

type Monitoring_MetricDescriptorLabel struct {
	// A human-readable description for the label.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The key for this label. The key must not exceed 100 characters. The first character of the key must be an upper- or lower-case letter, the remaining characters must be letters, digits or underscores, and the key must match the regular expression [a-zA-Z][a-zA-Z0-9_]-
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	/*
	   The type of data that can be assigned to the label.
	   Default value is `STRING`.
	   Possible values are: `STRING`, `BOOL`, `INT64`.
	*/
	ValueType string `json:"valueType,omitempty" yaml:"valueType,omitempty"`
}
