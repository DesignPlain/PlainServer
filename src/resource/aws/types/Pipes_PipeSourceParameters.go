package types

type Pipes_PipeSourceParameters struct {
	// The parameters for using a Amazon SQS stream as a source. Detailed below.
	SqsQueueParameters Pipes_PipeSourceParametersSqsQueueParameters `json:"sqsQueueParameters,omitempty" yaml:"sqsQueueParameters,omitempty"`

	// The parameters for using an Active MQ broker as a source. Detailed below.
	ActivemqBrokerParameters Pipes_PipeSourceParametersActivemqBrokerParameters `json:"activemqBrokerParameters,omitempty" yaml:"activemqBrokerParameters,omitempty"`

	// The parameters for using a DynamoDB stream as a source.  Detailed below.
	DynamodbStreamParameters Pipes_PipeSourceParametersDynamodbStreamParameters `json:"dynamodbStreamParameters,omitempty" yaml:"dynamodbStreamParameters,omitempty"`

	// The collection of event patterns used to [filter events](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-filtering.html). Detailed below.
	FilterCriteria Pipes_PipeSourceParametersFilterCriteria `json:"filterCriteria,omitempty" yaml:"filterCriteria,omitempty"`

	// The parameters for using a Kinesis stream as a source. Detailed below.
	KinesisStreamParameters Pipes_PipeSourceParametersKinesisStreamParameters `json:"kinesisStreamParameters,omitempty" yaml:"kinesisStreamParameters,omitempty"`

	// The parameters for using an MSK stream as a source. Detailed below.
	ManagedStreamingKafkaParameters Pipes_PipeSourceParametersManagedStreamingKafkaParameters `json:"managedStreamingKafkaParameters,omitempty" yaml:"managedStreamingKafkaParameters,omitempty"`

	// The parameters for using a Rabbit MQ broker as a source. Detailed below.
	RabbitmqBrokerParameters Pipes_PipeSourceParametersRabbitmqBrokerParameters `json:"rabbitmqBrokerParameters,omitempty" yaml:"rabbitmqBrokerParameters,omitempty"`

	// The parameters for using a self-managed Apache Kafka stream as a source. Detailed below.
	SelfManagedKafkaParameters Pipes_PipeSourceParametersSelfManagedKafkaParameters `json:"selfManagedKafkaParameters,omitempty" yaml:"selfManagedKafkaParameters,omitempty"`
}
