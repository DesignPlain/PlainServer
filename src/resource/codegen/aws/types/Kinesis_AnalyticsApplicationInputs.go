package types

type Kinesis_AnalyticsApplicationInputs struct {
	/*
	   The point at which the application starts processing records from the streaming source.
	   See Starting Position Configuration below for more details.
	*/
	StartingPositionConfigurations []Kinesis_AnalyticsApplicationInputsStartingPositionConfiguration `json:"startingPositionConfigurations,omitempty" yaml:"startingPositionConfigurations,omitempty"`

	/*
	   The Kinesis Firehose configuration for the streaming source. Conflicts with `kinesis_stream`.
	   See Kinesis Firehose below for more details.
	*/
	KinesisFirehose Kinesis_AnalyticsApplicationInputsKinesisFirehose `json:"kinesisFirehose,omitempty" yaml:"kinesisFirehose,omitempty"`

	/*
	   The number of Parallel in-application streams to create.
	   See Parallelism below for more details.
	*/
	Parallelism Kinesis_AnalyticsApplicationInputsParallelism `json:"parallelism,omitempty" yaml:"parallelism,omitempty"`

	/*
	   The Processing Configuration to transform records as they are received from the stream.
	   See Processing Configuration below for more details.
	*/
	ProcessingConfiguration Kinesis_AnalyticsApplicationInputsProcessingConfiguration `json:"processingConfiguration,omitempty" yaml:"processingConfiguration,omitempty"`

	// The Schema format of the data in the streaming source. See Source Schema below for more details.
	Schema Kinesis_AnalyticsApplicationInputsSchema `json:"schema,omitempty" yaml:"schema,omitempty"`

	// The ARN of the Kinesis Analytics Application.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	/*
	   The Kinesis Stream configuration for the streaming source. Conflicts with `kinesis_firehose`.
	   See Kinesis Stream below for more details.
	*/
	KinesisStream Kinesis_AnalyticsApplicationInputsKinesisStream `json:"kinesisStream,omitempty" yaml:"kinesisStream,omitempty"`

	// The Name Prefix to use when creating an in-application stream.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	//
	StreamNames []string `json:"streamNames,omitempty" yaml:"streamNames,omitempty"`
}
