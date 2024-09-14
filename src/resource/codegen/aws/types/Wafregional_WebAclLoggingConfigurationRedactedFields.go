package types

type Wafregional_WebAclLoggingConfigurationRedactedFields struct {
	// Set of configuration blocks for fields to redact. Detailed below.
	FieldToMatches []Wafregional_WebAclLoggingConfigurationRedactedFieldsFieldToMatch `json:"fieldToMatches,omitempty" yaml:"fieldToMatches,omitempty"`
}
