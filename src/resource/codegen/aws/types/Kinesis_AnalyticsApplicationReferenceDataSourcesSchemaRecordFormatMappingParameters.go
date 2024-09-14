package types

type Kinesis_AnalyticsApplicationReferenceDataSourcesSchemaRecordFormatMappingParameters struct {
	/*
	   Mapping information when the record format uses delimiters.
	   See CSV Mapping Parameters below for more details.
	*/
	Csv Kinesis_AnalyticsApplicationReferenceDataSourcesSchemaRecordFormatMappingParametersCsv `json:"csv,omitempty" yaml:"csv,omitempty"`

	/*
	   Mapping information when JSON is the record format on the streaming source.
	   See JSON Mapping Parameters below for more details.
	*/
	Json Kinesis_AnalyticsApplicationReferenceDataSourcesSchemaRecordFormatMappingParametersJson `json:"json,omitempty" yaml:"json,omitempty"`
}
