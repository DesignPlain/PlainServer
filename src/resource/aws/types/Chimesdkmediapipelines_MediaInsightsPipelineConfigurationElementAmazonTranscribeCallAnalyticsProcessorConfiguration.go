package types

type Chimesdkmediapipelines_MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfiguration struct {
	// Method for applying a vocabulary filter to Transcript events.
	VocabularyFilterMethod string `json:"vocabularyFilterMethod,omitempty" yaml:"vocabularyFilterMethod,omitempty"`

	// Language code for the transcription model.
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	// Name of custom language model for transcription.
	LanguageModelName string `json:"languageModelName,omitempty" yaml:"languageModelName,omitempty"`

	// Level of stability to use when partial results stabilization is enabled.
	PartialResultsStability string `json:"partialResultsStability,omitempty" yaml:"partialResultsStability,omitempty"`

	// Types of personally identifiable information (PII) to redact from a Transcript event.
	PiiEntityTypes string `json:"piiEntityTypes,omitempty" yaml:"piiEntityTypes,omitempty"`

	// Settings for post call analytics.
	PostCallAnalyticsSettings Chimesdkmediapipelines_MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationPostCallAnalyticsSettings `json:"postCallAnalyticsSettings,omitempty" yaml:"postCallAnalyticsSettings,omitempty"`

	// Name of the custom vocabulary filter to use when processing Transcript events.
	VocabularyFilterName string `json:"vocabularyFilterName,omitempty" yaml:"vocabularyFilterName,omitempty"`

	// Name of the custom vocabulary to use when processing Transcript events.
	VocabularyName string `json:"vocabularyName,omitempty" yaml:"vocabularyName,omitempty"`

	// Filter for category events to be delivered to insights target.
	CallAnalyticsStreamCategories []string `json:"callAnalyticsStreamCategories,omitempty" yaml:"callAnalyticsStreamCategories,omitempty"`

	// Labels all personally identifiable information (PII) identified in Transcript events.
	ContentIdentificationType string `json:"contentIdentificationType,omitempty" yaml:"contentIdentificationType,omitempty"`

	// Redacts all personally identifiable information (PII) identified in Transcript events.
	ContentRedactionType string `json:"contentRedactionType,omitempty" yaml:"contentRedactionType,omitempty"`

	// Enables partial result stabilization in Transcript events.
	EnablePartialResultsStabilization bool `json:"enablePartialResultsStabilization,omitempty" yaml:"enablePartialResultsStabilization,omitempty"`

	// Filters partial Utterance events from delivery to the insights target.
	FilterPartialResults bool `json:"filterPartialResults,omitempty" yaml:"filterPartialResults,omitempty"`
}
