package types

type Kinesis_AnalyticsApplicationInputsKinesisStream struct {
	// The ARN of the Kinesis Stream.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	// The ARN of the IAM Role used to access the stream.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
