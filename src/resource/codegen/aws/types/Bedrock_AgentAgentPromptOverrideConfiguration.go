package types

type Bedrock_AgentAgentPromptOverrideConfiguration struct {
	// ARN of the Lambda function to use when parsing the raw foundation model output in parts of the agent sequence. If you specify this field, at least one of the `prompt_configurations` block must contain a `parser_mode` value that is set to `OVERRIDDEN`.
	OverrideLambda string `json:"overrideLambda,omitempty" yaml:"overrideLambda,omitempty"`

	// Configurations to override a prompt template in one part of an agent sequence. See `prompt_configurations` Block for details.
	PromptConfigurations []Bedrock_AgentAgentPromptOverrideConfigurationPromptConfiguration `json:"promptConfigurations,omitempty" yaml:"promptConfigurations,omitempty"`
}
