package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesSnowflakeErrorHandlingConfig struct {
	// Name of the Amazon S3 bucket.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// Amazon S3 bucket prefix.
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`

	// If the flow should fail after the first instance of a failure when attempting to place data in the destination.
	FailOnFirstDestinationError bool `json:"failOnFirstDestinationError,omitempty" yaml:"failOnFirstDestinationError,omitempty"`
}
