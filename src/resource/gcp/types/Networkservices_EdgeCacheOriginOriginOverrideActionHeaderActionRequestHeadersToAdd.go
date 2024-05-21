package types

type Networkservices_EdgeCacheOriginOriginOverrideActionHeaderActionRequestHeadersToAdd struct {
	// The value of the header to add.
	HeaderValue string `json:"headerValue,omitempty" yaml:"headerValue,omitempty"`

	/*
	   Whether to replace all existing headers with the same name.
	   By default, added header values are appended
	   to the response or request headers with the
	   same field names. The added values are
	   separated by commas.
	   To overwrite existing values, set `replace` to `true`.
	*/
	Replace bool `json:"replace,omitempty" yaml:"replace,omitempty"`

	// The name of the header to add.
	HeaderName string `json:"headerName,omitempty" yaml:"headerName,omitempty"`
}
