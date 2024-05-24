package types

type Appmesh_VirtualNodeSpecListenerTlsValidationSubjectAlternativeNamesMatch struct {
	// Values sent must match the specified values exactly.
	Exacts []string `json:"exacts,omitempty" yaml:"exacts,omitempty"`
}
