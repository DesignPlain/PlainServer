package types

type Pipes_PipeSourceParametersManagedStreamingKafkaParameters struct {
	// The maximum number of records to include in each batch. Maximum value of 10000.
	BatchSize int `json:"batchSize,omitempty" yaml:"batchSize,omitempty"`

	// The name of the destination queue to consume. Maximum value of 200.
	ConsumerGroupId string `json:"consumerGroupId,omitempty" yaml:"consumerGroupId,omitempty"`

	// The credentials needed to access the resource. Detailed below.
	Credentials Pipes_PipeSourceParametersManagedStreamingKafkaParametersCredentials `json:"credentials,omitempty" yaml:"credentials,omitempty"`

	// The maximum length of a time to wait for events. Maximum value of 300.
	MaximumBatchingWindowInSeconds int `json:"maximumBatchingWindowInSeconds,omitempty" yaml:"maximumBatchingWindowInSeconds,omitempty"`

	// The position in a stream from which to start reading. Valid values: TRIM_HORIZON, LATEST.
	StartingPosition string `json:"startingPosition,omitempty" yaml:"startingPosition,omitempty"`

	// The name of the topic that the pipe will read from. Maximum length of 249.
	TopicName string `json:"topicName,omitempty" yaml:"topicName,omitempty"`
}
