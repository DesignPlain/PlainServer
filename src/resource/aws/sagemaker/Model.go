package sagemaker

import types "DesignSphere_Server/src/resource/aws/types"

type Model struct {
	// The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
	PrimaryContainer types.Sagemaker_ModelPrimaryContainer `json:"primaryContainer,omitempty" yaml:"primaryContainer,omitempty"`

	/*
	   A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.

	   The `primary_container` and `container` block both support:
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
	VpcConfig types.Sagemaker_ModelVpcConfig `json:"vpcConfig,omitempty" yaml:"vpcConfig,omitempty"`

	// Specifies containers in the inference pipeline. If not specified, the `primary_container` argument is required. Fields are documented below.
	Containers []types.Sagemaker_ModelContainer `json:"containers,omitempty" yaml:"containers,omitempty"`

	// Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
	EnableNetworkIsolation bool `json:"enableNetworkIsolation,omitempty" yaml:"enableNetworkIsolation,omitempty"`

	// A role that SageMaker can assume to access model artifacts and docker images for deployment.
	ExecutionRoleArn string `json:"executionRoleArn,omitempty" yaml:"executionRoleArn,omitempty"`

	// Specifies details of how containers in a multi-container endpoint are called. see Inference Execution Config.
	InferenceExecutionConfig types.Sagemaker_ModelInferenceExecutionConfig `json:"inferenceExecutionConfig,omitempty" yaml:"inferenceExecutionConfig,omitempty"`

	// The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
