package types

type Bedrock_GuardrailSensitiveInformationPolicyConfigPiiEntitiesConfig struct {
	// Options for sensitive information action.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// The currently supported PII entities.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
