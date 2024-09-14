package transcribe

type MedicalVocabulary struct {
	// The Amazon S3 location (URI) of the text file that contains your custom medical vocabulary.
	VocabularyFileUri string `json:"vocabularyFileUri,omitempty" yaml:"vocabularyFileUri,omitempty"`

	/*
	   The name of the Medical Vocabulary.

	   The following arguments are optional:
	*/
	VocabularyName string `json:"vocabularyName,omitempty" yaml:"vocabularyName,omitempty"`

	// The language code you selected for your medical vocabulary. US English (en-US) is the only language supported with Amazon Transcribe Medical.
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	// A map of tags to assign to the MedicalVocabulary. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
