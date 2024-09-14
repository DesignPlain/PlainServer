package types

type Sagemaker_FlowDefinitionHumanLoopActivationConfigHumanLoopActivationConditionsConfig struct {
	// A JSON expressing use-case specific conditions declaratively. If any condition is matched, atomic tasks are created against the configured work team. For more information about how to structure the JSON, see [JSON Schema for Human Loop Activation Conditions in Amazon Augmented AI](https://docs.aws.amazon.com/sagemaker/latest/dg/a2i-human-fallback-conditions-json-schema.html).
	HumanLoopActivationConditions string `json:"humanLoopActivationConditions,omitempty" yaml:"humanLoopActivationConditions,omitempty"`
}
