package types

type Bedrockmodel_InvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3Config struct {
	// S3 bucket name.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// S3 prefix.
	KeyPrefix string `json:"keyPrefix,omitempty" yaml:"keyPrefix,omitempty"`
}
