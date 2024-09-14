package types

type Appmesh_VirtualGatewaySpecListenerTlsValidationSubjectAlternativeNamesMatch struct {
	// Values sent must match the specified values exactly.
	Exacts []string `json:"exacts,omitempty" yaml:"exacts,omitempty"`
}
