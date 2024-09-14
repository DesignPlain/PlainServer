package types

type Bedrock_AgentAgentPromptOverrideConfigurationPromptConfigurationInferenceConfiguration struct {
	// Likelihood of the model selecting higher-probability options while generating a response. A lower value makes the model more likely to choose higher-probability options, while a higher value makes the model more likely to choose lower-probability options.
	Temperature float64 `json:"temperature,omitempty" yaml:"temperature,omitempty"`

	// Number of top most-likely candidates, between 0 and 500, from which the model chooses the next token in the sequence.
	TopK int `json:"topK,omitempty" yaml:"topK,omitempty"`

	// Top percentage of the probability distribution of next tokens, between 0 and 1 (denoting 0%!a(MISSING)nd 100%!)(MISSING), from which the model chooses the next token in the sequence.
	TopP float64 `json:"topP,omitempty" yaml:"topP,omitempty"`

	// Maximum number of tokens to allow in the generated response.
	MaxLength int `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`

	// List of stop sequences. A stop sequence is a sequence of characters that causes the model to stop generating the response.
	StopSequences []string `json:"stopSequences,omitempty" yaml:"stopSequences,omitempty"`
}
