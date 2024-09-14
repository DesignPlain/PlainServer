package types

type Kinesis_AnalyticsApplicationOutputSchema struct {
	// The Format Type of the records on the output stream. Can be `CSV` or `JSON`.
	RecordFormatType string `json:"recordFormatType,omitempty" yaml:"recordFormatType,omitempty"`
}
