package types

type Appmesh_VirtualGatewaySpecListenerTlsCertificate struct {
	// Local file certificate.
	File Appmesh_VirtualGatewaySpecListenerTlsCertificateFile `json:"file,omitempty" yaml:"file,omitempty"`

	// A [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.
	Sds Appmesh_VirtualGatewaySpecListenerTlsCertificateSds `json:"sds,omitempty" yaml:"sds,omitempty"`

	// An AWS Certificate Manager (ACM) certificate.
	Acm Appmesh_VirtualGatewaySpecListenerTlsCertificateAcm `json:"acm,omitempty" yaml:"acm,omitempty"`
}
