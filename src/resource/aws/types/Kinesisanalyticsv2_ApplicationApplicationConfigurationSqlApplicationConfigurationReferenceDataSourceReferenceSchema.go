package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchema struct {
	/*
	   Specifies the encoding of the records in the streaming source. For example, `UTF-8`.

	   The `s3_reference_data_source` object supports the following:
	*/
	RecordEncoding string `json:"recordEncoding,omitempty" yaml:"recordEncoding,omitempty"`

	// Specifies the format of the records on the streaming source.
	RecordFormat Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormat `json:"recordFormat,omitempty" yaml:"recordFormat,omitempty"`

	// Describes the mapping of each data element in the streaming source to the corresponding column in the in-application stream.
	RecordColumns []Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordColumn `json:"recordColumns,omitempty" yaml:"recordColumns,omitempty"`
}
