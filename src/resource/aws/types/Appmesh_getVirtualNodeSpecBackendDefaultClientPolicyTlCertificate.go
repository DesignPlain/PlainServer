package types

type Appmesh_getVirtualNodeSpecBackendDefaultClientPolicyTlCertificate struct {
	//
	Files []Appmesh_getVirtualNodeSpecBackendDefaultClientPolicyTlCertificateFile `json:"files,omitempty" yaml:"files,omitempty"`

	//
	Sds []Appmesh_getVirtualNodeSpecBackendDefaultClientPolicyTlCertificateSd `json:"sds,omitempty" yaml:"sds,omitempty"`
}
