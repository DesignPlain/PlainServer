package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationInput struct {
	// Describes the format of the data in the streaming source, and how each data element maps to corresponding columns in the in-application stream that is being created.
	InputSchema Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchema `json:"inputSchema,omitempty" yaml:"inputSchema,omitempty"`

	// The point at which the application starts processing records from the streaming source.
	InputStartingPositionConfigurations []Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfiguration `json:"inputStartingPositionConfigurations,omitempty" yaml:"inputStartingPositionConfigurations,omitempty"`

	/*
	   The input processing configuration for the input.
	   An input processor transforms records as they are received from the stream, before the application's SQL code executes.
	*/
	InputProcessingConfiguration Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputProcessingConfiguration `json:"inputProcessingConfiguration,omitempty" yaml:"inputProcessingConfiguration,omitempty"`

	//
	InputId string `json:"inputId,omitempty" yaml:"inputId,omitempty"`

	// Describes the number of in-application streams to create.
	InputParallelism Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputParallelism `json:"inputParallelism,omitempty" yaml:"inputParallelism,omitempty"`

	// If the streaming source is a Kinesis Data Firehose delivery stream, identifies the delivery stream's ARN.
	KinesisFirehoseInput Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisFirehoseInput `json:"kinesisFirehoseInput,omitempty" yaml:"kinesisFirehoseInput,omitempty"`

	// If the streaming source is a Kinesis data stream, identifies the stream's Amazon Resource Name (ARN).
	KinesisStreamsInput Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInput `json:"kinesisStreamsInput,omitempty" yaml:"kinesisStreamsInput,omitempty"`

	// The name prefix to use when creating an in-application stream.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	//
	InAppStreamNames []string `json:"inAppStreamNames,omitempty" yaml:"inAppStreamNames,omitempty"`
}
