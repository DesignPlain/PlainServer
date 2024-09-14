package types

type Sagemaker_EndpointConfigurationAsyncInferenceConfigOutputConfig struct {
	// The Amazon Web Services Key Management Service (Amazon Web Services KMS) key that Amazon SageMaker uses to encrypt the asynchronous inference output in Amazon S3.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Specifies the configuration for notifications of inference results for asynchronous inference.
	NotificationConfig Sagemaker_EndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfig `json:"notificationConfig,omitempty" yaml:"notificationConfig,omitempty"`

	// The Amazon S3 location to upload failure inference responses to.
	S3FailurePath string `json:"s3FailurePath,omitempty" yaml:"s3FailurePath,omitempty"`

	// The Amazon S3 location to upload inference responses to.
	S3OutputPath string `json:"s3OutputPath,omitempty" yaml:"s3OutputPath,omitempty"`
}
