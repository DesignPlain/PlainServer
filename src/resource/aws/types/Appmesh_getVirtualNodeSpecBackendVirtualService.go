package types

type Appmesh_getVirtualNodeSpecBackendVirtualService struct {
	//
	ClientPolicies []Appmesh_getVirtualNodeSpecBackendVirtualServiceClientPolicy `json:"clientPolicies,omitempty" yaml:"clientPolicies,omitempty"`

	//
	VirtualServiceName string `json:"virtualServiceName,omitempty" yaml:"virtualServiceName,omitempty"`
}
