package types

type Appmesh_VirtualNodeSpecListenerTls struct {
	// Listener's TLS certificate.
	Certificate Appmesh_VirtualNodeSpecListenerTlsCertificate `json:"certificate,omitempty" yaml:"certificate,omitempty"`

	// Listener's TLS mode. Valid values: `DISABLED`, `PERMISSIVE`, `STRICT`.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	// Listener's Transport Layer Security (TLS) validation context.
	Validation Appmesh_VirtualNodeSpecListenerTlsValidation `json:"validation,omitempty" yaml:"validation,omitempty"`
}
