package types

type Chimesdkmediapipelines_MediaInsightsPipelineConfigurationElementAmazonTranscribeProcessorConfiguration struct {
	// Labels all personally identifiable information (PII) identified in Transcript events.
	ContentIdentificationType string `json:"contentIdentificationType,omitempty" yaml:"contentIdentificationType,omitempty"`

	// Name of custom language model for transcription.
	LanguageModelName string `json:"languageModelName,omitempty" yaml:"languageModelName,omitempty"`

	// Types of personally identifiable information (PII) to redact from a Transcript event.
	PiiEntityTypes string `json:"piiEntityTypes,omitempty" yaml:"piiEntityTypes,omitempty"`

	// Enables speaker partitioning (diarization) in your Transcript events.
	ShowSpeakerLabel bool `json:"showSpeakerLabel,omitempty" yaml:"showSpeakerLabel,omitempty"`

	// Name of the custom vocabulary filter to use when processing Transcript events.
	VocabularyFilterName string `json:"vocabularyFilterName,omitempty" yaml:"vocabularyFilterName,omitempty"`

	// Redacts all personally identifiable information (PII) identified in Transcript events.
	ContentRedactionType string `json:"contentRedactionType,omitempty" yaml:"contentRedactionType,omitempty"`

	// Enables partial result stabilization in Transcript events.
	EnablePartialResultsStabilization bool `json:"enablePartialResultsStabilization,omitempty" yaml:"enablePartialResultsStabilization,omitempty"`

	// Filters partial Utterance events from delivery to the insights target.
	FilterPartialResults bool `json:"filterPartialResults,omitempty" yaml:"filterPartialResults,omitempty"`

	// Language code for the transcription model.
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	// Level of stability to use when partial results stabilization is enabled.
	PartialResultsStability string `json:"partialResultsStability,omitempty" yaml:"partialResultsStability,omitempty"`

	// Method for applying a vocabulary filter to Transcript events.
	VocabularyFilterMethod string `json:"vocabularyFilterMethod,omitempty" yaml:"vocabularyFilterMethod,omitempty"`

	// Name of the custom vocabulary to use when processing Transcript events.
	VocabularyName string `json:"vocabularyName,omitempty" yaml:"vocabularyName,omitempty"`
}
