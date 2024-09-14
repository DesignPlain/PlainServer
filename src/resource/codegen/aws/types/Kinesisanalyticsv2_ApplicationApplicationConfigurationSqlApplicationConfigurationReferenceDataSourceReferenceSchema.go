package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchema struct {
	// Describes the mapping of each data element in the streaming source to the corresponding column in the in-application stream.
	RecordColumns []Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordColumn `json:"recordColumns,omitempty" yaml:"recordColumns,omitempty"`

	// Specifies the encoding of the records in the streaming source. For example, `UTF-8`.
	RecordEncoding string `json:"recordEncoding,omitempty" yaml:"recordEncoding,omitempty"`

	// Specifies the format of the records on the streaming source.
	RecordFormat Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormat `json:"recordFormat,omitempty" yaml:"recordFormat,omitempty"`
}
