package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationOutput struct {
	// Identifies a Kinesis Data Firehose delivery stream as the destination.
	KinesisFirehoseOutput Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationOutputKinesisFirehoseOutput `json:"kinesisFirehoseOutput,omitempty" yaml:"kinesisFirehoseOutput,omitempty"`

	// Identifies a Kinesis data stream as the destination.
	KinesisStreamsOutput Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationOutputKinesisStreamsOutput `json:"kinesisStreamsOutput,omitempty" yaml:"kinesisStreamsOutput,omitempty"`

	// Identifies a Lambda function as the destination.
	LambdaOutput Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationOutputLambdaOutput `json:"lambdaOutput,omitempty" yaml:"lambdaOutput,omitempty"`

	// The name of the in-application stream.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	OutputId string `json:"outputId,omitempty" yaml:"outputId,omitempty"`

	// Describes the data format when records are written to the destination.
	DestinationSchema Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchema `json:"destinationSchema,omitempty" yaml:"destinationSchema,omitempty"`
}
