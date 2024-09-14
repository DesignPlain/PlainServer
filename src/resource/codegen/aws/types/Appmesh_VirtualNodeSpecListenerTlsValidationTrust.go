package types

type Appmesh_VirtualNodeSpecListenerTlsValidationTrust struct {
	// TLS validation context trust for a local file certificate.
	File Appmesh_VirtualNodeSpecListenerTlsValidationTrustFile `json:"file,omitempty" yaml:"file,omitempty"`

	// TLS validation context trust for a [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.
	Sds Appmesh_VirtualNodeSpecListenerTlsValidationTrustSds `json:"sds,omitempty" yaml:"sds,omitempty"`
}
