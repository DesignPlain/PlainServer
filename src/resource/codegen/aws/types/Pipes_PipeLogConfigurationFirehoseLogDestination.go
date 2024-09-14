package types

type Pipes_PipeLogConfigurationFirehoseLogDestination struct {
	// Amazon Resource Name (ARN) of the Kinesis Data Firehose delivery stream to which EventBridge delivers the pipe log records.
	DeliveryStreamArn string `json:"deliveryStreamArn,omitempty" yaml:"deliveryStreamArn,omitempty"`
}
