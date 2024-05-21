package types

type Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleDictionary struct {
	/*
	   Newline-delimited file of words in Cloud Storage. Only a single file is accepted.
	   Structure is documented below.
	*/
	CloudStoragePath Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleDictionaryCloudStoragePath `json:"cloudStoragePath,omitempty" yaml:"cloudStoragePath,omitempty"`

	/*
	   List of words or phrases to search for.
	   Structure is documented below.
	*/
	WordList Dataloss_PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleDictionaryWordList `json:"wordList,omitempty" yaml:"wordList,omitempty"`
}
