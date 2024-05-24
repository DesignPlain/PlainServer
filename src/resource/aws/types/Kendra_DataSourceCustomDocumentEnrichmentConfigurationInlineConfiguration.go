package types

type Kendra_DataSourceCustomDocumentEnrichmentConfigurationInlineConfiguration struct {
	// Configuration of the condition used for the target document attribute or metadata field when ingesting documents into Amazon Kendra. See condition.
	Condition Kendra_DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	// `TRUE` to delete content if the condition used for the target attribute is met.
	DocumentContentDeletion bool `json:"documentContentDeletion,omitempty" yaml:"documentContentDeletion,omitempty"`

	// Configuration of the target document attribute or metadata field when ingesting documents into Amazon Kendra. You can also include a value. Detailed below.
	Target Kendra_DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTarget `json:"target,omitempty" yaml:"target,omitempty"`
}
