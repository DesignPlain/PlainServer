package types

type Appmesh_VirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrust struct {
	// TLS validation context trust for an AWS Certificate Manager (ACM) certificate.
	Acm Appmesh_VirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustAcm `json:"acm,omitempty" yaml:"acm,omitempty"`

	// TLS validation context trust for a local file certificate.
	File Appmesh_VirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustFile `json:"file,omitempty" yaml:"file,omitempty"`

	// TLS validation context trust for a [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.
	Sds Appmesh_VirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustSds `json:"sds,omitempty" yaml:"sds,omitempty"`
}
