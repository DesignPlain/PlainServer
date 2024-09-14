package types

type Kinesis_FirehoseDeliveryStreamMskSourceConfiguration struct {
	// The authentication configuration of the Amazon MSK cluster. See `authentication_configuration` block below for details.
	AuthenticationConfiguration Kinesis_FirehoseDeliveryStreamMskSourceConfigurationAuthenticationConfiguration `json:"authenticationConfiguration,omitempty" yaml:"authenticationConfiguration,omitempty"`

	// The ARN of the Amazon MSK cluster.
	MskClusterArn string `json:"mskClusterArn,omitempty" yaml:"mskClusterArn,omitempty"`

	// The topic name within the Amazon MSK cluster.
	TopicName string `json:"topicName,omitempty" yaml:"topicName,omitempty"`
}
