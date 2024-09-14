package types

type Bedrock_AgentAgentPromptOverrideConfigurationPromptConfiguration struct {
	// Whether to override the default parser Lambda function when parsing the raw foundation model output in the part of the agent sequence defined by the `prompt_type`. If you set the argument as `OVERRIDDEN`, the `override_lambda` argument in the `prompt_override_configuration` block must be specified with the ARN of a Lambda function. Valid values: `DEFAULT`, `OVERRIDDEN`.
	ParserMode string `json:"parserMode,omitempty" yaml:"parserMode,omitempty"`

	// Whether to override the default prompt template for this `prompt_type`. Set this argument to `OVERRIDDEN` to use the prompt that you provide in the `base_prompt_template`. If you leave it as `DEFAULT`, the agent uses a default prompt template. Valid values: `DEFAULT`, `OVERRIDDEN`.
	PromptCreationMode string `json:"promptCreationMode,omitempty" yaml:"promptCreationMode,omitempty"`

	// Whether to allow the agent to carry out the step specified in the `prompt_type`. If you set this argument to `DISABLED`, the agent skips that step. Valid Values: `ENABLED`, `DISABLED`.
	PromptState string `json:"promptState,omitempty" yaml:"promptState,omitempty"`

	// Step in the agent sequence that this prompt configuration applies to. Valid values: `PRE_PROCESSING`, `ORCHESTRATION`, `POST_PROCESSING`, `KNOWLEDGE_BASE_RESPONSE_GENERATION`.
	PromptType string `json:"promptType,omitempty" yaml:"promptType,omitempty"`

	// prompt template with which to replace the default prompt template. You can use placeholder variables in the base prompt template to customize the prompt. For more information, see [Prompt template placeholder variables](https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-placeholders.html).
	BasePromptTemplate string `json:"basePromptTemplate,omitempty" yaml:"basePromptTemplate,omitempty"`

	// Inference parameters to use when the agent invokes a foundation model in the part of the agent sequence defined by the `prompt_type`. For more information, see [Inference parameters for foundation models](https://docs.aws.amazon.com/bedrock/latest/userguide/model-parameters.html). See `inference_configuration` Block for details.
	InferenceConfigurations []Bedrock_AgentAgentPromptOverrideConfigurationPromptConfigurationInferenceConfiguration `json:"inferenceConfigurations,omitempty" yaml:"inferenceConfigurations,omitempty"`
}
