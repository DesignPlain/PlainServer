package types

type Sagemaker_DataQualityJobDefinitionDataQualityJobInput struct {
	// Input object for the endpoint. Fields are documented below.
	EndpointInput Sagemaker_DataQualityJobDefinitionDataQualityJobInputEndpointInput `json:"endpointInput,omitempty" yaml:"endpointInput,omitempty"`

	// Input object for the batch transform job. Fields are documented below.
	BatchTransformInput Sagemaker_DataQualityJobDefinitionDataQualityJobInputBatchTransformInput `json:"batchTransformInput,omitempty" yaml:"batchTransformInput,omitempty"`
}
