package types

type Iot_TopicRuleErrorActionKafka struct {
	// The list of Kafka headers that you specify. Nested arguments below.
	Headers []Iot_TopicRuleErrorActionKafkaHeader `json:"headers,omitempty" yaml:"headers,omitempty"`

	// The Kafka message key.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The Kafka message partition.
	Partition string `json:"partition,omitempty" yaml:"partition,omitempty"`

	// The Kafka topic for messages to be sent to the Kafka broker.
	Topic string `json:"topic,omitempty" yaml:"topic,omitempty"`

	// Properties of the Apache Kafka producer client. For more info, see the [AWS documentation](https://docs.aws.amazon.com/iot/latest/developerguide/apache-kafka-rule-action.html).
	ClientProperties map[string]string `json:"clientProperties,omitempty" yaml:"clientProperties,omitempty"`

	// The ARN of Kafka action's VPC `aws.iot.TopicRuleDestination`.
	DestinationArn string `json:"destinationArn,omitempty" yaml:"destinationArn,omitempty"`
}
