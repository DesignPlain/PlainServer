package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesS3 struct {
	// Name of the Amazon S3 bucket.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// Amazon S3 bucket prefix.
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`

	// Configuration that determines how Amazon AppFlow should format the flow output data when Amazon S3 is used as the destination. See S3 Output Format Config for more details.
	S3OutputFormatConfig Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfig `json:"s3OutputFormatConfig,omitempty" yaml:"s3OutputFormatConfig,omitempty"`
}
