package types

type Waf_WebAclLoggingConfigurationRedactedFieldsFieldToMatch struct {
	// When the value of `type` is `HEADER`, enter the name of the header that you want the WAF to search, for example, `User-Agent` or `Referer`. If the value of `type` is any other value, omit `data`.
	Data string `json:"data,omitempty" yaml:"data,omitempty"`

	// The part of the web request that you want AWS WAF to search for a specified stringE.g., `HEADER` or `METHOD`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
