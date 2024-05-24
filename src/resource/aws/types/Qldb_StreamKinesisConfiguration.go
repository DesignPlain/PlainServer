package types

type Qldb_StreamKinesisConfiguration struct {
	// The Amazon Resource Name (ARN) of the Kinesis Data Streams resource.
	StreamArn string `json:"streamArn,omitempty" yaml:"streamArn,omitempty"`

	// Enables QLDB to publish multiple data records in a single Kinesis Data Streams record, increasing the number of records sent per API call. Default: `true`.
	AggregationEnabled bool `json:"aggregationEnabled,omitempty" yaml:"aggregationEnabled,omitempty"`
}
