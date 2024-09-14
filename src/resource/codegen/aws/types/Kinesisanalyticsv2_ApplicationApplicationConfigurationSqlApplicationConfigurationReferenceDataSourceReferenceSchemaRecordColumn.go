package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordColumn struct {
	// A reference to the data element in the streaming input or the reference data source.
	Mapping string `json:"mapping,omitempty" yaml:"mapping,omitempty"`

	// The name of the column that is created in the in-application input stream or reference table.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The type of column created in the in-application input stream or reference table.
	SqlType string `json:"sqlType,omitempty" yaml:"sqlType,omitempty"`
}
