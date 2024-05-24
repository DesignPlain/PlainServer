package transcribe

type VocabularyFilter struct {
	// The Amazon S3 location (URI) of the text file that contains your custom VocabularyFilter. Conflicts with `words` argument.
	VocabularyFilterFileUri string `json:"vocabularyFilterFileUri,omitempty" yaml:"vocabularyFilterFileUri,omitempty"`

	/*
	   The name of the VocabularyFilter.

	   The following arguments are optional:
	*/
	VocabularyFilterName string `json:"vocabularyFilterName,omitempty" yaml:"vocabularyFilterName,omitempty"`

	// A list of terms to include in the vocabulary. Conflicts with `vocabulary_filter_file_uri` argument.
	Words []string `json:"words,omitempty" yaml:"words,omitempty"`

	// The language code you selected for your vocabulary filter. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	// A map of tags to assign to the VocabularyFilter. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
