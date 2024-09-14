package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersCsvMappingParameters struct {
	// The column delimiter. For example, in a CSV format, a comma (`,`) is the typical column delimiter.
	RecordColumnDelimiter string `json:"recordColumnDelimiter,omitempty" yaml:"recordColumnDelimiter,omitempty"`

	// The row delimiter. For example, in a CSV format, `\n` is the typical row delimiter.
	RecordRowDelimiter string `json:"recordRowDelimiter,omitempty" yaml:"recordRowDelimiter,omitempty"`
}
