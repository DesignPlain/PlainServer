package types

type Sagemaker_ModelInferenceExecutionConfig struct {
	// The container hosts value `SingleModel/MultiModel`. The default value is `SingleModel`.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`
}
