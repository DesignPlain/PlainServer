package types

type Sesv2_ConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestination struct {
	// The Amazon Resource Name (ARN) of the Amazon Kinesis Data Firehose stream that the Amazon SES API v2 sends email events to.
	DeliveryStreamArn string `json:"deliveryStreamArn,omitempty" yaml:"deliveryStreamArn,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role that the Amazon SES API v2 uses to send email events to the Amazon Kinesis Data Firehose stream.
	IamRoleArn string `json:"iamRoleArn,omitempty" yaml:"iamRoleArn,omitempty"`
}
