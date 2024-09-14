package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormat struct {
	// Provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
	MappingParameters Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParameters `json:"mappingParameters,omitempty" yaml:"mappingParameters,omitempty"`

	// The type of record format. Valid values: `CSV`, `JSON`.
	RecordFormatType string `json:"recordFormatType,omitempty" yaml:"recordFormatType,omitempty"`
}
