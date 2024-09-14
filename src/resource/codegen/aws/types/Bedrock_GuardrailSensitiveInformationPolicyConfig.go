package types

type Bedrock_GuardrailSensitiveInformationPolicyConfig struct {
	// List of entities. See PII Entities Config for more information.
	PiiEntitiesConfigs []Bedrock_GuardrailSensitiveInformationPolicyConfigPiiEntitiesConfig `json:"piiEntitiesConfigs,omitempty" yaml:"piiEntitiesConfigs,omitempty"`

	// List of regex. See Regexes Config for more information.
	RegexesConfigs []Bedrock_GuardrailSensitiveInformationPolicyConfigRegexesConfig `json:"regexesConfigs,omitempty" yaml:"regexesConfigs,omitempty"`
}
