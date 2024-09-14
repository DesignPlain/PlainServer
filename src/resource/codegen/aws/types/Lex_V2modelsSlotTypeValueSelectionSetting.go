package types

type Lex_V2modelsSlotTypeValueSelectionSetting struct {
	// Provides settings that enable advanced recognition settings for slot values. You can use this to enable using slot values as a custom vocabulary for recognizing user utterances. See [`advanced_recognition_setting` argument reference] below.
	AdvancedRecognitionSettings []Lex_V2modelsSlotTypeValueSelectionSettingAdvancedRecognitionSetting `json:"advancedRecognitionSettings,omitempty" yaml:"advancedRecognitionSettings,omitempty"`

	// Used to validate the value of the slot. See [`regex_filter` argument reference] below.
	RegexFilters []Lex_V2modelsSlotTypeValueSelectionSettingRegexFilter `json:"regexFilters,omitempty" yaml:"regexFilters,omitempty"`

	// Determines the slot resolution strategy that Amazon Lex uses to return slot type values. The field can be set to one of the following values: `ORIGINAL_VALUE` - Returns the value entered by the user, if the user value is similar to the slot value. `TOP_RESOLUTION` If there is a resolution list for the slot, return the first value in the resolution list as the slot type value. If there is no resolution list, null is returned. If you don't specify the valueSelectionStrategy , the default is `ORIGINAL_VALUE`. Valid values are `OriginalValue`, `TopResolution`, and `Concatenation`.
	ResolutionStrategy string `json:"resolutionStrategy,omitempty" yaml:"resolutionStrategy,omitempty"`
}
