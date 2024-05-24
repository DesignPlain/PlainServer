package types

type Appmesh_VirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificate struct {
	// A [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.
	Sds Appmesh_VirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificateSds `json:"sds,omitempty" yaml:"sds,omitempty"`

	// Local file certificate.
	File Appmesh_VirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificateFile `json:"file,omitempty" yaml:"file,omitempty"`
}
