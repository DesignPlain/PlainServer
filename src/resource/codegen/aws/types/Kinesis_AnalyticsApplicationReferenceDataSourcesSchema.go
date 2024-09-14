package types

type Kinesis_AnalyticsApplicationReferenceDataSourcesSchema struct {
	/*
	   The Record Column mapping for the streaming source data element.
	   See Record Columns below for more details.
	*/
	RecordColumns []Kinesis_AnalyticsApplicationReferenceDataSourcesSchemaRecordColumn `json:"recordColumns,omitempty" yaml:"recordColumns,omitempty"`

	// The Encoding of the record in the streaming source.
	RecordEncoding string `json:"recordEncoding,omitempty" yaml:"recordEncoding,omitempty"`

	/*
	   The Record Format and mapping information to schematize a record.
	   See Record Format below for more details.
	*/
	RecordFormat Kinesis_AnalyticsApplicationReferenceDataSourcesSchemaRecordFormat `json:"recordFormat,omitempty" yaml:"recordFormat,omitempty"`
}
