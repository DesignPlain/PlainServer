package types

type Bedrockmodel_InvocationLoggingConfigurationLoggingConfigCloudwatchConfig struct {
	// S3 configuration for delivering a large amount of data.
	LargeDataDeliveryS3Config Bedrockmodel_InvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3Config `json:"largeDataDeliveryS3Config,omitempty" yaml:"largeDataDeliveryS3Config,omitempty"`

	// Log group name.
	LogGroupName string `json:"logGroupName,omitempty" yaml:"logGroupName,omitempty"`

	// The role ARN.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
