package types

type Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrust struct {
	// TLS validation context trust for an AWS Certificate Manager (ACM) certificate.
	Acm Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcm `json:"acm,omitempty" yaml:"acm,omitempty"`

	// TLS validation context trust for a local file certificate.
	File Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustFile `json:"file,omitempty" yaml:"file,omitempty"`

	// TLS validation context trust for a [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.
	Sds Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustSds `json:"sds,omitempty" yaml:"sds,omitempty"`
}
