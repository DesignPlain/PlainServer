package types

type Chimesdkmediapipelines_MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfiguration struct {
	// Redacts all personally identifiable information (PII) identified in Utterance events.
	ContentRedactionType string `json:"contentRedactionType,omitempty" yaml:"contentRedactionType,omitempty"`

	// Filters partial Utterance events from delivery to the insights target.
	FilterPartialResults bool `json:"filterPartialResults,omitempty" yaml:"filterPartialResults,omitempty"`

	// Types of personally identifiable information (PII) to redact from an Utterance event.
	PiiEntityTypes string `json:"piiEntityTypes,omitempty" yaml:"piiEntityTypes,omitempty"`

	// Settings for post call analytics.
	PostCallAnalyticsSettings Chimesdkmediapipelines_MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationPostCallAnalyticsSettings `json:"postCallAnalyticsSettings,omitempty" yaml:"postCallAnalyticsSettings,omitempty"`

	// Name of the custom vocabulary filter to use when processing Utterance events.
	VocabularyFilterName string `json:"vocabularyFilterName,omitempty" yaml:"vocabularyFilterName,omitempty"`

	// Name of the custom vocabulary to use when processing Utterance events.
	VocabularyName string `json:"vocabularyName,omitempty" yaml:"vocabularyName,omitempty"`

	// Filter for category events to be delivered to insights target.
	CallAnalyticsStreamCategories []string `json:"callAnalyticsStreamCategories,omitempty" yaml:"callAnalyticsStreamCategories,omitempty"`

	// Enables partial result stabilization in Utterance events.
	EnablePartialResultsStabilization bool `json:"enablePartialResultsStabilization,omitempty" yaml:"enablePartialResultsStabilization,omitempty"`

	// Language code for the transcription model.
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	// Name of custom language model for transcription.
	LanguageModelName string `json:"languageModelName,omitempty" yaml:"languageModelName,omitempty"`

	// Level of stability to use when partial results stabilization is enabled.
	PartialResultsStability string `json:"partialResultsStability,omitempty" yaml:"partialResultsStability,omitempty"`

	// Method for applying a vocabulary filter to Utterance events.
	VocabularyFilterMethod string `json:"vocabularyFilterMethod,omitempty" yaml:"vocabularyFilterMethod,omitempty"`

	// Labels all personally identifiable information (PII) identified in Utterance events.
	ContentIdentificationType string `json:"contentIdentificationType,omitempty" yaml:"contentIdentificationType,omitempty"`
}
