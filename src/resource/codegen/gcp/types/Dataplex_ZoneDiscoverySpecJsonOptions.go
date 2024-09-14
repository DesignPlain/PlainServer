package types

type Dataplex_ZoneDiscoverySpecJsonOptions struct {
	// Optional. Whether to disable the inference of data type for Json data. If true, all columns will be registered as their primitive types (strings, number or boolean).
	DisableTypeInference bool `json:"disableTypeInference,omitempty" yaml:"disableTypeInference,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	Encoding string `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}
