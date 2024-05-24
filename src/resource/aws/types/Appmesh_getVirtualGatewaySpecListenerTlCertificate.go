package types

type Appmesh_getVirtualGatewaySpecListenerTlCertificate struct {
	//
	Acms []Appmesh_getVirtualGatewaySpecListenerTlCertificateAcm `json:"acms,omitempty" yaml:"acms,omitempty"`

	//
	Files []Appmesh_getVirtualGatewaySpecListenerTlCertificateFile `json:"files,omitempty" yaml:"files,omitempty"`

	//
	Sds []Appmesh_getVirtualGatewaySpecListenerTlCertificateSd `json:"sds,omitempty" yaml:"sds,omitempty"`
}
