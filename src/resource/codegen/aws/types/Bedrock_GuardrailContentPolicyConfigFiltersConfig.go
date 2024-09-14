package types

type Bedrock_GuardrailContentPolicyConfigFiltersConfig struct {
	// Strength for filters.
	InputStrength string `json:"inputStrength,omitempty" yaml:"inputStrength,omitempty"`

	// Strength for filters.
	OutputStrength string `json:"outputStrength,omitempty" yaml:"outputStrength,omitempty"`

	// Type of contextual grounding filter.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
