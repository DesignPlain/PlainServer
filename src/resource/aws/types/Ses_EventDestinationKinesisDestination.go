package types

type Ses_EventDestinationKinesisDestination struct {
	// The ARN of the role that has permissions to access the Kinesis Stream
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The ARN of the Kinesis Stream
	StreamArn string `json:"streamArn,omitempty" yaml:"streamArn,omitempty"`
}
