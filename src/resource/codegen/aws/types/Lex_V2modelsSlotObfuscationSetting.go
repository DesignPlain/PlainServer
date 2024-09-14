package types

type Lex_V2modelsSlotObfuscationSetting struct {
	// Whether Amazon Lex obscures slot values in conversation logs. Valid values are `DefaultObfuscation` and `None`.
	ObfuscationSettingType string `json:"obfuscationSettingType,omitempty" yaml:"obfuscationSettingType,omitempty"`
}
