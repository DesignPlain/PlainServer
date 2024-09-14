package types

type Kinesis_AnalyticsApplicationReferenceDataSourcesSchemaRecordColumn struct {
	// Name of the column.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The SQL Type of the column.
	SqlType string `json:"sqlType,omitempty" yaml:"sqlType,omitempty"`

	// The Mapping reference to the data element.
	Mapping string `json:"mapping,omitempty" yaml:"mapping,omitempty"`
}
