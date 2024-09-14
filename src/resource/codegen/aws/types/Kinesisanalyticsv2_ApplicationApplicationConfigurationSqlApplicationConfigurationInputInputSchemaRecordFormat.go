package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormat struct {
	// The type of record format. Valid values: `CSV`, `JSON`.
	RecordFormatType string `json:"recordFormatType,omitempty" yaml:"recordFormatType,omitempty"`

	// Provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
	MappingParameters Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParameters `json:"mappingParameters,omitempty" yaml:"mappingParameters,omitempty"`
}
