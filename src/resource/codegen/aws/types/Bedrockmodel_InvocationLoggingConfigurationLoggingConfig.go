package types

type Bedrockmodel_InvocationLoggingConfigurationLoggingConfig struct {
	// CloudWatch logging configuration.
	CloudwatchConfig Bedrockmodel_InvocationLoggingConfigurationLoggingConfigCloudwatchConfig `json:"cloudwatchConfig,omitempty" yaml:"cloudwatchConfig,omitempty"`

	// Set to include embeddings data in the log delivery.
	EmbeddingDataDeliveryEnabled bool `json:"embeddingDataDeliveryEnabled,omitempty" yaml:"embeddingDataDeliveryEnabled,omitempty"`

	// Set to include image data in the log delivery.
	ImageDataDeliveryEnabled bool `json:"imageDataDeliveryEnabled,omitempty" yaml:"imageDataDeliveryEnabled,omitempty"`

	// S3 configuration for storing log data.
	S3Config Bedrockmodel_InvocationLoggingConfigurationLoggingConfigS3Config `json:"s3Config,omitempty" yaml:"s3Config,omitempty"`

	// Set to include text data in the log delivery.
	TextDataDeliveryEnabled bool `json:"textDataDeliveryEnabled,omitempty" yaml:"textDataDeliveryEnabled,omitempty"`
}
