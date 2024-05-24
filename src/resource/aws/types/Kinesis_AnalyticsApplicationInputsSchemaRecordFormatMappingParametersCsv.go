package types

type Kinesis_AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsv struct {
	// The Column Delimiter.
	RecordColumnDelimiter string `json:"recordColumnDelimiter,omitempty" yaml:"recordColumnDelimiter,omitempty"`

	// The Row Delimiter.
	RecordRowDelimiter string `json:"recordRowDelimiter,omitempty" yaml:"recordRowDelimiter,omitempty"`
}
