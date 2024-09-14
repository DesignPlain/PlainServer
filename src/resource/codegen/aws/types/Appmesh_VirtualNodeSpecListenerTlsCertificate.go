package types

type Appmesh_VirtualNodeSpecListenerTlsCertificate struct {
	// An AWS Certificate Manager (ACM) certificate.
	Acm Appmesh_VirtualNodeSpecListenerTlsCertificateAcm `json:"acm,omitempty" yaml:"acm,omitempty"`

	// Local file certificate.
	File Appmesh_VirtualNodeSpecListenerTlsCertificateFile `json:"file,omitempty" yaml:"file,omitempty"`

	// A [Secret Discovery Service](https://www.envoyproxy.io/docs/envoy/latest/configuration/security/secret#secret-discovery-service-sds) certificate.
	Sds Appmesh_VirtualNodeSpecListenerTlsCertificateSds `json:"sds,omitempty" yaml:"sds,omitempty"`
}
