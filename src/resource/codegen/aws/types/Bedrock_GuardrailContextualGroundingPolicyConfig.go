package types

type Bedrock_GuardrailContextualGroundingPolicyConfig struct {
	// List of contextual grounding filter configs. See Contextual Grounding Filters Config for more information.
	FiltersConfigs []Bedrock_GuardrailContextualGroundingPolicyConfigFiltersConfig `json:"filtersConfigs,omitempty" yaml:"filtersConfigs,omitempty"`
}
