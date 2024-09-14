package types

type Kinesis_AnalyticsApplicationReferenceDataSourcesSchemaRecordFormat struct {
	/*
	   The Mapping Information for the record format.
	   See Mapping Parameters below for more details.
	*/
	MappingParameters Kinesis_AnalyticsApplicationReferenceDataSourcesSchemaRecordFormatMappingParameters `json:"mappingParameters,omitempty" yaml:"mappingParameters,omitempty"`

	// The type of Record Format. Can be `CSV` or `JSON`.
	RecordFormatType string `json:"recordFormatType,omitempty" yaml:"recordFormatType,omitempty"`
}
