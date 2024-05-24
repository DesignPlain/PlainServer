package types

type Sagemaker_FlowDefinitionHumanLoopActivationConfig struct {
	// defines under what conditions SageMaker creates a human loop. See Human Loop Activation Conditions Config details below.
	HumanLoopActivationConditionsConfig Sagemaker_FlowDefinitionHumanLoopActivationConfigHumanLoopActivationConditionsConfig `json:"humanLoopActivationConditionsConfig,omitempty" yaml:"humanLoopActivationConditionsConfig,omitempty"`
}
