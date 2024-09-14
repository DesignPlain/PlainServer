package types

type Appmesh_getVirtualServiceSpecProvider struct {
	//
	VirtualNodes []Appmesh_getVirtualServiceSpecProviderVirtualNode `json:"virtualNodes,omitempty" yaml:"virtualNodes,omitempty"`

	//
	VirtualRouters []Appmesh_getVirtualServiceSpecProviderVirtualRouter `json:"virtualRouters,omitempty" yaml:"virtualRouters,omitempty"`
}
