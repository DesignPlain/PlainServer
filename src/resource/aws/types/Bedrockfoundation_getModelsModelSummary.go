package types

type Bedrockfoundation_getModelsModelSummary struct {
	// Model provider name.
	ProviderName string `json:"providerName,omitempty" yaml:"providerName,omitempty"`

	// Indicates whether the model supports streaming.
	ResponseStreamingSupported bool `json:"responseStreamingSupported,omitempty" yaml:"responseStreamingSupported,omitempty"`

	// Customizations that the model supports.
	CustomizationsSupporteds []string `json:"customizationsSupporteds,omitempty" yaml:"customizationsSupporteds,omitempty"`

	// Inference types that the model supports.
	InferenceTypesSupporteds []string `json:"inferenceTypesSupporteds,omitempty" yaml:"inferenceTypesSupporteds,omitempty"`

	// Model ARN.
	ModelArn string `json:"modelArn,omitempty" yaml:"modelArn,omitempty"`

	// Output modalities that the model supports.
	OutputModalities []string `json:"outputModalities,omitempty" yaml:"outputModalities,omitempty"`

	// Input modalities that the model supports.
	InputModalities []string `json:"inputModalities,omitempty" yaml:"inputModalities,omitempty"`

	// Model identifier.
	ModelId string `json:"modelId,omitempty" yaml:"modelId,omitempty"`

	// Model name.
	ModelName string `json:"modelName,omitempty" yaml:"modelName,omitempty"`
}
