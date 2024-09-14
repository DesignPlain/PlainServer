package types

type Kendra_DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTarget struct {
	// The identifier of the target document attribute or metadata field. For example, 'Department' could be an identifier for the target attribute or metadata field that includes the department names associated with the documents.
	TargetDocumentAttributeKey string `json:"targetDocumentAttributeKey,omitempty" yaml:"targetDocumentAttributeKey,omitempty"`

	// The target value you want to create for the target attribute. For example, 'Finance' could be the target value for the target attribute key 'Department'. See target_document_attribute_value.
	TargetDocumentAttributeValue Kendra_DataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationTargetTargetDocumentAttributeValue `json:"targetDocumentAttributeValue,omitempty" yaml:"targetDocumentAttributeValue,omitempty"`

	// `TRUE` to delete the existing target value for your specified target attribute key. You cannot create a target value and set this to `TRUE`. To create a target value (`TargetDocumentAttributeValue`), set this to `FALSE`.
	TargetDocumentAttributeValueDeletion bool `json:"targetDocumentAttributeValueDeletion,omitempty" yaml:"targetDocumentAttributeValueDeletion,omitempty"`
}
