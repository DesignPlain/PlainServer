package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesS3 struct {
	//
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	//
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`

	//
	S3OutputFormatConfig Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfig `json:"s3OutputFormatConfig,omitempty" yaml:"s3OutputFormatConfig,omitempty"`
}
