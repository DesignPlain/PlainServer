package transcribe

import types "DesignSphere_Server/src/resource/aws/types"

type LanguageModel struct {
	// Name of reference base model.
	BaseModelName string `json:"baseModelName,omitempty" yaml:"baseModelName,omitempty"`

	// The input data config for the LanguageModel. See Input Data Config for more details.
	InputDataConfig types.Transcribe_LanguageModelInputDataConfig `json:"inputDataConfig,omitempty" yaml:"inputDataConfig,omitempty"`

	// The language code you selected for your language model. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	// The model name.
	ModelName string `json:"modelName,omitempty" yaml:"modelName,omitempty"`

	// A map of tags to assign to the LanguageModel. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
