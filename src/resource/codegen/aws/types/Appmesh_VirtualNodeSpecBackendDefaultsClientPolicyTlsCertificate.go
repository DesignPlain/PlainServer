package types

type Appmesh_VirtualNodeSpecBackendDefaultsClientPolicyTlsCertificate struct {
	// Local file certificate.
	File Appmesh_VirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFile `json:"file,omitempty" yaml:"file,omitempty"`

	// A [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.
	Sds Appmesh_VirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateSds `json:"sds,omitempty" yaml:"sds,omitempty"`
}
