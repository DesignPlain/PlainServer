package types

type Pipes_PipeSourceParametersSelfManagedKafkaParameters struct {
	// The name of the destination queue to consume. Maximum value of 200.
	ConsumerGroupId string `json:"consumerGroupId,omitempty" yaml:"consumerGroupId,omitempty"`

	// The ARN of the Secrets Manager secret used for certification.
	ServerRootCaCertificate string `json:"serverRootCaCertificate,omitempty" yaml:"serverRootCaCertificate,omitempty"`

	// The position in a stream from which to start reading. Valid values: TRIM_HORIZON, LATEST.
	StartingPosition string `json:"startingPosition,omitempty" yaml:"startingPosition,omitempty"`

	// This structure specifies the VPC subnets and security groups for the stream, and whether a public IP address is to be used. Detailed below.
	Vpc Pipes_PipeSourceParametersSelfManagedKafkaParametersVpc `json:"vpc,omitempty" yaml:"vpc,omitempty"`

	// An array of server URLs. Maximum number of 2 items, each of maximum length 300.
	AdditionalBootstrapServers []string `json:"additionalBootstrapServers,omitempty" yaml:"additionalBootstrapServers,omitempty"`

	// The credentials needed to access the resource. Detailed below.
	Credentials Pipes_PipeSourceParametersSelfManagedKafkaParametersCredentials `json:"credentials,omitempty" yaml:"credentials,omitempty"`

	// The maximum length of a time to wait for events. Maximum value of 300.
	MaximumBatchingWindowInSeconds int `json:"maximumBatchingWindowInSeconds,omitempty" yaml:"maximumBatchingWindowInSeconds,omitempty"`

	// The name of the topic that the pipe will read from. Maximum length of 249.
	TopicName string `json:"topicName,omitempty" yaml:"topicName,omitempty"`

	// The maximum number of records to include in each batch. Maximum value of 10000.
	BatchSize int `json:"batchSize,omitempty" yaml:"batchSize,omitempty"`
}
