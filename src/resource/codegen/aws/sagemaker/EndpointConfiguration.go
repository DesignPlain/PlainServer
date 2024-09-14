package sagemaker

import types "libds/aws/types"

type EndpointConfiguration struct {
	// Creates a unique endpoint configuration name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// An list of ProductionVariant objects, one for each model that you want to host at this endpoint. Fields are documented below.
	ProductionVariants []types.Sagemaker_EndpointConfigurationProductionVariant `json:"productionVariants,omitempty" yaml:"productionVariants,omitempty"`

	// Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants. If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants. Fields are documented below.
	ShadowProductionVariants []types.Sagemaker_EndpointConfigurationShadowProductionVariant `json:"shadowProductionVariants,omitempty" yaml:"shadowProductionVariants,omitempty"`

	// A mapping of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies configuration for how an endpoint performs asynchronous inference.
	AsyncInferenceConfig types.Sagemaker_EndpointConfigurationAsyncInferenceConfig `json:"asyncInferenceConfig,omitempty" yaml:"asyncInferenceConfig,omitempty"`

	// Specifies the parameters to capture input/output of SageMaker models endpoints. Fields are documented below.
	DataCaptureConfig types.Sagemaker_EndpointConfigurationDataCaptureConfig `json:"dataCaptureConfig,omitempty" yaml:"dataCaptureConfig,omitempty"`

	// Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// The name of the endpoint configuration. If omitted, this provider will assign a random, unique name. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
