package types

type Endpoints_ServiceApiMethod struct {
	// The simple name of the endpoint as described in the config.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The type URL for the request to this API.
	RequestType string `json:"requestType,omitempty" yaml:"requestType,omitempty"`

	// The type URL for the response from this API.
	ResponseType string `json:"responseType,omitempty" yaml:"responseType,omitempty"`

	// `SYNTAX_PROTO2` or `SYNTAX_PROTO3`.
	Syntax string `json:"syntax,omitempty" yaml:"syntax,omitempty"`
}
