package types

type Appmesh_getVirtualNodeSpecListenerTlCertificate struct {
	//
	Files []Appmesh_getVirtualNodeSpecListenerTlCertificateFile `json:"files,omitempty" yaml:"files,omitempty"`

	//
	Sds []Appmesh_getVirtualNodeSpecListenerTlCertificateSd `json:"sds,omitempty" yaml:"sds,omitempty"`

	//
	Acms []Appmesh_getVirtualNodeSpecListenerTlCertificateAcm `json:"acms,omitempty" yaml:"acms,omitempty"`
}
