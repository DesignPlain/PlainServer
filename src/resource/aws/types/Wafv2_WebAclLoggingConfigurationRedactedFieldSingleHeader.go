package types

type Wafv2_WebAclLoggingConfigurationRedactedFieldSingleHeader struct {
	// Name of the query header to redact. This setting must be provided in lowercase characters.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
