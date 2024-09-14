package types

type Kinesis_AnalyticsApplicationOutput struct {
	// The ARN of the Kinesis Analytics Application.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	/*
	   The Kinesis Firehose configuration for the destination stream. Conflicts with `kinesis_stream`.
	   See Kinesis Firehose below for more details.
	*/
	KinesisFirehose Kinesis_AnalyticsApplicationOutputKinesisFirehose `json:"kinesisFirehose,omitempty" yaml:"kinesisFirehose,omitempty"`

	/*
	   The Kinesis Stream configuration for the destination stream. Conflicts with `kinesis_firehose`.
	   See Kinesis Stream below for more details.
	*/
	KinesisStream Kinesis_AnalyticsApplicationOutputKinesisStream `json:"kinesisStream,omitempty" yaml:"kinesisStream,omitempty"`

	// The Lambda function destination. See Lambda below for more details.
	Lambda Kinesis_AnalyticsApplicationOutputLambda `json:"lambda,omitempty" yaml:"lambda,omitempty"`

	// The Name of the in-application stream.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The Schema format of the data written to the destination. See Destination Schema below for more details.
	Schema Kinesis_AnalyticsApplicationOutputSchema `json:"schema,omitempty" yaml:"schema,omitempty"`
}
