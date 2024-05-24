package types

type Appmesh_getVirtualGatewaySpecBackendDefaultClientPolicyTlCertificate struct {
	//
	Sds []Appmesh_getVirtualGatewaySpecBackendDefaultClientPolicyTlCertificateSd `json:"sds,omitempty" yaml:"sds,omitempty"`

	//
	Files []Appmesh_getVirtualGatewaySpecBackendDefaultClientPolicyTlCertificateFile `json:"files,omitempty" yaml:"files,omitempty"`
}
