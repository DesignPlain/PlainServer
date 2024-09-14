package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesSnowflake struct {
	//
	ErrorHandlingConfig Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesSnowflakeErrorHandlingConfig `json:"errorHandlingConfig,omitempty" yaml:"errorHandlingConfig,omitempty"`

	//
	IntermediateBucketName string `json:"intermediateBucketName,omitempty" yaml:"intermediateBucketName,omitempty"`

	//
	Object string `json:"object,omitempty" yaml:"object,omitempty"`

	//
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`
}
