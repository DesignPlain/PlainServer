package types

type Pipes_PipeSourceParametersActivemqBrokerParameters struct {
	// The maximum number of records to include in each batch. Maximum value of 10000.
	BatchSize int `json:"batchSize,omitempty" yaml:"batchSize,omitempty"`

	// The credentials needed to access the resource. Detailed below.
	Credentials Pipes_PipeSourceParametersActivemqBrokerParametersCredentials `json:"credentials,omitempty" yaml:"credentials,omitempty"`

	// The maximum length of a time to wait for events. Maximum value of 300.
	MaximumBatchingWindowInSeconds int `json:"maximumBatchingWindowInSeconds,omitempty" yaml:"maximumBatchingWindowInSeconds,omitempty"`

	// The name of the destination queue to consume. Maximum length of 1000.
	QueueName string `json:"queueName,omitempty" yaml:"queueName,omitempty"`
}
