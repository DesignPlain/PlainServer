package types

type Wafv2_WebAclLoggingConfigurationRedactedField struct {
	// Configuration block that redacts the request URI path. It should be specified as an empty configuration block `{}`. The URI path is the part of a web request that identifies a resource, such as `/images/daily-ad.jpg`.
	UriPath Wafv2_WebAclLoggingConfigurationRedactedFieldUriPath `json:"uriPath,omitempty" yaml:"uriPath,omitempty"`

	// HTTP method to be redacted. It must be specified as an empty configuration block `{}`. The method indicates the type of operation that the request is asking the origin to perform.
	Method Wafv2_WebAclLoggingConfigurationRedactedFieldMethod `json:"method,omitempty" yaml:"method,omitempty"`

	// Whether to redact the query string. It must be specified as an empty configuration block `{}`. The query string is the part of a URL that appears after a `?` character, if any.
	QueryString Wafv2_WebAclLoggingConfigurationRedactedFieldQueryString `json:"queryString,omitempty" yaml:"queryString,omitempty"`

	// "single_header" refers to the redaction of a single header. For more information, please see the details below under Single Header.
	SingleHeader Wafv2_WebAclLoggingConfigurationRedactedFieldSingleHeader `json:"singleHeader,omitempty" yaml:"singleHeader,omitempty"`
}
