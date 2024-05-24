package transcribe

type Vocabulary struct {
	/*
	   The name of the Vocabulary.

	   The following arguments are optional:
	*/
	VocabularyName string `json:"vocabularyName,omitempty" yaml:"vocabularyName,omitempty"`

	// The language code you selected for your vocabulary.
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	// A list of terms to include in the vocabulary. Conflicts with `vocabulary_file_uri`
	Phrases []string `json:"phrases,omitempty" yaml:"phrases,omitempty"`

	// A map of tags to assign to the Vocabulary. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The Amazon S3 location (URI) of the text file that contains your custom vocabulary. Conflicts wth `phrases`.
	VocabularyFileUri string `json:"vocabularyFileUri,omitempty" yaml:"vocabularyFileUri,omitempty"`
}
