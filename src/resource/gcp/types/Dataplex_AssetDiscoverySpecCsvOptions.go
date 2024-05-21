package types

type Dataplex_AssetDiscoverySpecCsvOptions struct {
	// Optional. The delimiter being used to separate values. This defaults to ','.
	Delimiter string `json:"delimiter,omitempty" yaml:"delimiter,omitempty"`

	// Optional. Whether to disable the inference of data type for CSV data. If true, all columns will be registered as strings.
	DisableTypeInference bool `json:"disableTypeInference,omitempty" yaml:"disableTypeInference,omitempty"`

	// Optional. The character encoding of the data. The default is UTF-8.
	Encoding string `json:"encoding,omitempty" yaml:"encoding,omitempty"`

	// Optional. The number of rows to interpret as header rows that should be skipped when reading data rows.
	HeaderRows int `json:"headerRows,omitempty" yaml:"headerRows,omitempty"`
}
