package types

type Bedrock_GuardrailWordPolicyConfig struct {
	// A config for the list of managed words. See Managed Word Lists Config for more information.
	ManagedWordListsConfigs []Bedrock_GuardrailWordPolicyConfigManagedWordListsConfig `json:"managedWordListsConfigs,omitempty" yaml:"managedWordListsConfigs,omitempty"`

	// List of custom word configs. See Words Config for more information.
	WordsConfigs []Bedrock_GuardrailWordPolicyConfigWordsConfig `json:"wordsConfigs,omitempty" yaml:"wordsConfigs,omitempty"`
}
