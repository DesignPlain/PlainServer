package types

type Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificate struct {
	// Local file certificate.
	File Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateFile `json:"file,omitempty" yaml:"file,omitempty"`

	// A [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.
	Sds Appmesh_VirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateSds `json:"sds,omitempty" yaml:"sds,omitempty"`
}
