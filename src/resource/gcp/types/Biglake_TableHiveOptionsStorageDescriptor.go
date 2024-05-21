package types

type Biglake_TableHiveOptionsStorageDescriptor struct {
	// The fully qualified Java class name of the input format.
	InputFormat string `json:"inputFormat,omitempty" yaml:"inputFormat,omitempty"`

	// Cloud Storage folder URI where the table data is stored, starting with "gs://".
	LocationUri string `json:"locationUri,omitempty" yaml:"locationUri,omitempty"`

	// The fully qualified Java class name of the output format.
	OutputFormat string `json:"outputFormat,omitempty" yaml:"outputFormat,omitempty"`
}
