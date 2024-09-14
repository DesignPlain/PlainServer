package types

type Dataloss_PreventionInspectTemplateInspectConfigCustomInfoTypeDictionary struct {
	/*
	   Newline-delimited file of words in Cloud Storage. Only a single file is accepted.
	   Structure is documented below.
	*/
	CloudStoragePath Dataloss_PreventionInspectTemplateInspectConfigCustomInfoTypeDictionaryCloudStoragePath `json:"cloudStoragePath,omitempty" yaml:"cloudStoragePath,omitempty"`

	/*
	   List of words or phrases to search for.
	   Structure is documented below.
	*/
	WordList Dataloss_PreventionInspectTemplateInspectConfigCustomInfoTypeDictionaryWordList `json:"wordList,omitempty" yaml:"wordList,omitempty"`
}
