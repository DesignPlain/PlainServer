package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesRedshift struct {
	//
	Object string `json:"object,omitempty" yaml:"object,omitempty"`

	//
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`

	//
	ErrorHandlingConfig Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesRedshiftErrorHandlingConfig `json:"errorHandlingConfig,omitempty" yaml:"errorHandlingConfig,omitempty"`

	//
	IntermediateBucketName string `json:"intermediateBucketName,omitempty" yaml:"intermediateBucketName,omitempty"`
}
