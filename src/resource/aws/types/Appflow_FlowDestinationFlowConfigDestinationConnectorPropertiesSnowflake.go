package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesSnowflake struct {
	// Amazon S3 bucket prefix.
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`

	// Settings that determine how Amazon AppFlow handles an error when placing data in the custom connector as destination. See Error Handling Config for more details.
	ErrorHandlingConfig Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesSnowflakeErrorHandlingConfig `json:"errorHandlingConfig,omitempty" yaml:"errorHandlingConfig,omitempty"`

	// Intermediate bucket that Amazon AppFlow uses when moving data into Amazon Redshift.
	IntermediateBucketName string `json:"intermediateBucketName,omitempty" yaml:"intermediateBucketName,omitempty"`

	// Object specified in the flow destination.
	Object string `json:"object,omitempty" yaml:"object,omitempty"`
}
