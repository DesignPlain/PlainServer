package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchema struct {
	// Specifies the format of the records on the output stream. Valid values: `CSV`, `JSON`.
	RecordFormatType string `json:"recordFormatType,omitempty" yaml:"recordFormatType,omitempty"`
}
