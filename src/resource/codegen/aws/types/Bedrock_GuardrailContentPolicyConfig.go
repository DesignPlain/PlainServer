package types

type Bedrock_GuardrailContentPolicyConfig struct {
	// List of content filter configs in content policy. See Filters Config for more information.
	FiltersConfigs []Bedrock_GuardrailContentPolicyConfigFiltersConfig `json:"filtersConfigs,omitempty" yaml:"filtersConfigs,omitempty"`
}
