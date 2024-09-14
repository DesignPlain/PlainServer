package pinpoint

type EventStream struct {
	// The IAM role that authorizes Amazon Pinpoint to publish events to the stream in your account.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The application ID.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`

	// The Amazon Resource Name (ARN) of the Amazon Kinesis stream or Firehose delivery stream to which you want to publish events.
	DestinationStreamArn string `json:"destinationStreamArn,omitempty" yaml:"destinationStreamArn,omitempty"`
}
