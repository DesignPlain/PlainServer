package types

type Lambda_EventSourceMappingSelfManagedKafkaEventSourceConfig struct {
	// A Kafka consumer group ID between 1 and 200 characters for use when creating this event source mapping. If one is not specified, this value will be automatically generated. See [SelfManagedKafkaEventSourceConfig Syntax](https://docs.aws.amazon.com/lambda/latest/dg/API_SelfManagedKafkaEventSourceConfig.html).
	ConsumerGroupId string `json:"consumerGroupId,omitempty" yaml:"consumerGroupId,omitempty"`
}
