package types

type Appmesh_getVirtualGatewaySpecBackendDefaultClientPolicyTlCertificate struct {
	//
	Files []Appmesh_getVirtualGatewaySpecBackendDefaultClientPolicyTlCertificateFile `json:"files,omitempty" yaml:"files,omitempty"`

	//
	Sds []Appmesh_getVirtualGatewaySpecBackendDefaultClientPolicyTlCertificateSd `json:"sds,omitempty" yaml:"sds,omitempty"`
}
