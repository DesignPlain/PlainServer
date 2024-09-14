package types

type Bedrock_getCustomModelsModelSummary struct {
	// Creation time of the model.
	CreationTime string `json:"creationTime,omitempty" yaml:"creationTime,omitempty"`

	// The ARN of the custom model.
	ModelArn string `json:"modelArn,omitempty" yaml:"modelArn,omitempty"`

	// The name of the custom model.
	ModelName string `json:"modelName,omitempty" yaml:"modelName,omitempty"`
}
