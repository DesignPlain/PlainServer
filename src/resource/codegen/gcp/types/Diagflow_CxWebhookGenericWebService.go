package types

type Diagflow_CxWebhookGenericWebService struct {
	// The HTTP request headers to send together with webhook requests.
	RequestHeaders map[string]string `json:"requestHeaders,omitempty" yaml:"requestHeaders,omitempty"`

	// Whether to use speech adaptation for speech recognition.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`

	// Specifies a list of allowed custom CA certificates (in DER format) for HTTPS verification.
	AllowedCaCerts []string `json:"allowedCaCerts,omitempty" yaml:"allowedCaCerts,omitempty"`
}
