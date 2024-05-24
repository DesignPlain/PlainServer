package types

type Iot_TopicRuleFirehose struct {
	// A character separator that is used to separate records written to the Firehose stream. Valid values are: '\n' (newline), '\t' (tab), '\r\n' (Windows newline), ',' (comma).
	Separator string `json:"separator,omitempty" yaml:"separator,omitempty"`

	// The payload that contains a JSON array of records will be sent to Kinesis Firehose via a batch call.
	BatchMode bool `json:"batchMode,omitempty" yaml:"batchMode,omitempty"`

	// The delivery stream name.
	DeliveryStreamName string `json:"deliveryStreamName,omitempty" yaml:"deliveryStreamName,omitempty"`

	// The IAM role ARN that grants access to the Amazon Kinesis Firehose stream.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
