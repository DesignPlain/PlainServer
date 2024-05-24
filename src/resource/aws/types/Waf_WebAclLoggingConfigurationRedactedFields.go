package types

type Waf_WebAclLoggingConfigurationRedactedFields struct {
	// Set of configuration blocks for fields to redact. Detailed below.
	FieldToMatches []Waf_WebAclLoggingConfigurationRedactedFieldsFieldToMatch `json:"fieldToMatches,omitempty" yaml:"fieldToMatches,omitempty"`
}
