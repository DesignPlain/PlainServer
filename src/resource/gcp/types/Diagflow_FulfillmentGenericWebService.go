package types

type Diagflow_FulfillmentGenericWebService struct {
	// The password for HTTP Basic authentication.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// The HTTP request headers to send together with fulfillment requests.
	RequestHeaders map[string]string `json:"requestHeaders,omitempty" yaml:"requestHeaders,omitempty"`

	// The fulfillment URI for receiving POST requests. It must use https protocol.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`

	// The user name for HTTP Basic authentication.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
