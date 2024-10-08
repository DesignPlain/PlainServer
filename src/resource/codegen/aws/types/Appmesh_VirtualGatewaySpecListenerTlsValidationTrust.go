package types

type Appmesh_VirtualGatewaySpecListenerTlsValidationTrust struct {
	// TLS validation context trust for a local file certificate.
	File Appmesh_VirtualGatewaySpecListenerTlsValidationTrustFile `json:"file,omitempty" yaml:"file,omitempty"`

	// TLS validation context trust for a [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.
	Sds Appmesh_VirtualGatewaySpecListenerTlsValidationTrustSds `json:"sds,omitempty" yaml:"sds,omitempty"`
}
