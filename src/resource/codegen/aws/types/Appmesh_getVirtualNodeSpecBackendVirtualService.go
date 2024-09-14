package types

type Appmesh_getVirtualNodeSpecBackendVirtualService struct {
	//
	VirtualServiceName string `json:"virtualServiceName,omitempty" yaml:"virtualServiceName,omitempty"`

	//
	ClientPolicies []Appmesh_getVirtualNodeSpecBackendVirtualServiceClientPolicy `json:"clientPolicies,omitempty" yaml:"clientPolicies,omitempty"`
}
