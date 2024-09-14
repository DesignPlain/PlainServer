package types

type Kendra_DataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationCondition struct {
	// The value used by the operator. For example, you can specify the value 'financial' for strings in the `_source_uri` field that partially match or contain this value. See condition_on_value.
	ConditionOnValue Kendra_DataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValue `json:"conditionOnValue,omitempty" yaml:"conditionOnValue,omitempty"`

	// The condition operator. For example, you can use `Contains` to partially match a string. Valid Values: `GreaterThan` | `GreaterThanOrEquals` | `LessThan` | `LessThanOrEquals` | `Equals` | `NotEquals` | `Contains` | `NotContains` | `Exists` | `NotExists` | `BeginsWith`.
	Operator string `json:"operator,omitempty" yaml:"operator,omitempty"`

	// The identifier of the document attribute used for the condition. For example, `_source_uri` could be an identifier for the attribute or metadata field that contains source URIs associated with the documents. Amazon Kendra currently does not support `_document_body` as an attribute key used for the condition.
	ConditionDocumentAttributeKey string `json:"conditionDocumentAttributeKey,omitempty" yaml:"conditionDocumentAttributeKey,omitempty"`
}
