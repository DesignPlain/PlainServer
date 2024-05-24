package types

type Ivschat_LoggingConfigurationDestinationConfigurationFirehose struct {
	// Name of the Amazon Kinesis Firehose delivery stream where chat activity will be logged.
	DeliveryStreamName string `json:"deliveryStreamName,omitempty" yaml:"deliveryStreamName,omitempty"`
}
