package types

type Kinesis_FirehoseDeliveryStreamKinesisSourceConfiguration struct {
	// The ARN of the role that provides access to the source Kinesis stream.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The kinesis stream used as the source of the firehose delivery stream.
	KinesisStreamArn string `json:"kinesisStreamArn,omitempty" yaml:"kinesisStreamArn,omitempty"`
}
