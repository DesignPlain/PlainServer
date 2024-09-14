package types

type Kendra_DataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationConditionConditionOnValue struct {
	// A date expressed as an ISO 8601 string. It is important for the time zone to be included in the ISO 8601 date-time format. As of this writing only UTC is supported. For example, `2012-03-25T12:30:10+00:00`.
	DateValue string `json:"dateValue,omitempty" yaml:"dateValue,omitempty"`

	// A long integer value.
	LongValue int `json:"longValue,omitempty" yaml:"longValue,omitempty"`

	// A list of strings.
	StringListValues []string `json:"stringListValues,omitempty" yaml:"stringListValues,omitempty"`

	//
	StringValue string `json:"stringValue,omitempty" yaml:"stringValue,omitempty"`
}
