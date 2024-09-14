package types

type Bedrock_GuardrailContextualGroundingPolicyConfigFiltersConfig struct {
	// The threshold for this filter.
	Threshold float64 `json:"threshold,omitempty" yaml:"threshold,omitempty"`

	// Type of contextual grounding filter.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
