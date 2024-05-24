package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataSuccessResponseHandlingConfig struct {
	// Name of the Amazon S3 bucket.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// Amazon S3 bucket prefix.
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`
}
