package types

type Sesv2_ConfigurationSetEventDestinationEventDestinationSnsDestination struct {
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to publish email events to.
	TopicArn string `json:"topicArn,omitempty" yaml:"topicArn,omitempty"`
}
