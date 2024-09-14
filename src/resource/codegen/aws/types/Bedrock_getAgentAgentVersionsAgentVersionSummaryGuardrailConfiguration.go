package types

type Bedrock_getAgentAgentVersionsAgentVersionSummaryGuardrailConfiguration struct {
	// Unique identifier of the guardrail.
	GuardrailIdentifier string `json:"guardrailIdentifier,omitempty" yaml:"guardrailIdentifier,omitempty"`

	// Version of the guardrail.
	GuardrailVersion string `json:"guardrailVersion,omitempty" yaml:"guardrailVersion,omitempty"`
}
