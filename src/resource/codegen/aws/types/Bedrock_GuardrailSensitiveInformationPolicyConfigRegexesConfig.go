package types

type Bedrock_GuardrailSensitiveInformationPolicyConfigRegexesConfig struct {
	// Options for sensitive information action.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// The regex description.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The regex name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The regex pattern.
	Pattern string `json:"pattern,omitempty" yaml:"pattern,omitempty"`
}
