package types

type Lambda_EventSourceMappingAmazonManagedKafkaEventSourceConfig struct {
	// A Kafka consumer group ID between 1 and 200 characters for use when creating this event source mapping. If one is not specified, this value will be automatically generated. See [AmazonManagedKafkaEventSourceConfig Syntax](https://docs.aws.amazon.com/lambda/latest/dg/API_AmazonManagedKafkaEventSourceConfig.html).
	ConsumerGroupId string `json:"consumerGroupId,omitempty" yaml:"consumerGroupId,omitempty"`
}
